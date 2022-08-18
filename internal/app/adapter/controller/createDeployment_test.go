package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/adapter/repository"
	repo_test "github.com/doyyan/kubernetes/internal/app/adapter/repository/mocks"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/doyyan/kubernetes/internal/app/domain/domainrepo"
	domainrepotest "github.com/doyyan/kubernetes/internal/app/domain/domainrepo/mocks"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	kube "k8s.io/client-go/kubernetes"
)

func Test_CreateDeployment(t *testing.T) {
	tests := map[string]struct {
		err     error
		dep     repository.Deployment
		testDep model.Deployment
	}{
		"pass createDeploymentSuccess": {
			err:     nil,
			dep:     saveDeploymentReturnsSuccess(),
			testDep: DeploymentData(),
		},
		"pass createDeploymentFail": {
			err:     errors.New(" DB save failure"),
			dep:     saveDeploymentReturnsFail(),
			testDep: DeploymentData(),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			router := setupRouter(test.dep)
			w := httptest.NewRecorder()
			data, _ := json.Marshal(test.testDep)
			req, _ := http.NewRequest(http.MethodPost, "/deployment", bytes.NewBuffer(data))
			router.ServeHTTP(w, req)
			/*			cont, _ := gin.CreateTestContext(w)
						controller.deploymentRepository = test.dep
						controller.createDeployment(cont)*/
			assert.Equal(t, 200, w.Code)

			var got gin.H
			err := json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, test.err, err)
			err = test.dep.Create(context.Background(), logrus.New(), domain.Deployment{})
			if test.err != nil {
				assert.Equal(t, test.err, err)
			}
		})
	}
}

func saveDeploymentReturnsFail() repository.Deployment {
	k8smock := repo_test.K8SMock{
		CreateDeploymentFunc: func(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error {
			return errors.New(" DB save failure")
		},
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := repository.Deployment{K8S: &k8smock, DBconn: SetDB()}
	return dep
}

func saveDeploymentReturnsSuccess() repository.Deployment {
	k8smock := repo_test.K8SMock{
		CreateDeploymentFunc: func(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error {
			return nil
		},
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := repository.Deployment{K8S: &k8smock, DBconn: SetDB()}
	return dep
}

func SetDB() *gorm.DB {
	db1, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db1,
		PreferSimpleProtocol: true,
	})
	db2, _ := gorm.Open(dialector, &gorm.Config{})
	return db2
}

func CreateController() Controller {
	return Controller{
		Context: context.Background(),
		Logger:  logrus.New(),
	}
}

func setupRouter(dep repository.Deployment) *gin.Engine {
	r := gin.Default()
	cont := CreateController()
	cont.deploymentRepository = createDeployment()
	r.POST("/deployment", cont.createDeployment)
	return r
}

func DeploymentData() model.Deployment {
	return model.Deployment{
		Name:          "testDeployment",
		Kind:          "deployment",
		Image:         "nginx:1.12",
		ContainerPort: 80,
		ContainerName: "",
		NameSpace:     "default",
		Labels:        map[string]string{"app": "demo"},
		Replicas:      10,
	}
}

func createDeployment() domainrepo.IDeploymentRepo {
	depRepoMock := domainrepotest.IDeploymentRepoMock{CreateFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
		return nil
	}}
	return &depRepoMock
}
