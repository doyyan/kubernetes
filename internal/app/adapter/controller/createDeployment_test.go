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
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/doyyan/kubernetes/internal/app/domain/domainrepo"
	domainrepotest "github.com/doyyan/kubernetes/internal/app/domain/domainrepo/mocks"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_CreateDeployment(t *testing.T) {
	tests := map[string]struct {
		err            error
		testDep        model.Deployment
		testcase       domainrepo.IDeploymentRepo
		output         string
		httpReturnCode int
	}{
		"pass createDeploymentSuccess": {
			err:            nil,
			testDep:        DeploymentData(),
			testcase:       createDeploymentSuccess(),
			output:         "deployment created",
			httpReturnCode: 200,
		},
		"pass createDeploymentFail": {
			err:            errors.New(" DB save failure"),
			testcase:       createDeploymentFail(),
			testDep:        DeploymentData(),
			httpReturnCode: 500,
			output:         "",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			router := setupRouter(test.testcase)
			w := httptest.NewRecorder()
			data, _ := json.Marshal(test.testDep)
			req, _ := http.NewRequest(http.MethodPost, "/deployment", bytes.NewBuffer(data))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.httpReturnCode, w.Code)
			if w.Code == 200 {
				var output map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &output)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, test.output, output["message"])
			}
		})
	}
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

func setupRouter(testcase domainrepo.IDeploymentRepo) *gin.Engine {
	r := gin.Default()
	cont := CreateController()
	cont.deploymentRepository = testcase
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

func createDeploymentSuccess() domainrepo.IDeploymentRepo {
	depRepoMock := domainrepotest.IDeploymentRepoMock{CreateFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
		return nil
	}}
	return &depRepoMock
}

func createDeploymentFail() domainrepo.IDeploymentRepo {
	depRepoMock := domainrepotest.IDeploymentRepoMock{CreateFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
		return errors.New(" DB save failure")
	}}
	return &depRepoMock
}
