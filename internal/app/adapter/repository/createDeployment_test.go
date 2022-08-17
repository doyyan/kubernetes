package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repo_test "github.com/doyyan/kubernetes/internal/app/adapter/repository/mocks"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	kube "k8s.io/client-go/kubernetes"
)

func Test_CreateDeployment(t *testing.T) {
	tests := map[string]struct {
		err error
		dep Deployment
	}{
		"pass createDeploymentSuccess": {
			err: nil,
			dep: saveDeploymentReturnsSuccess(),
		},
		"pass createDeploymentFail": {
			err: errors.New(" DB save failure"),
			dep: saveDeploymentReturnsFail(),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.dep.Create(context.Background(), logrus.New(), domain.Deployment{})
			if test.err != nil {
				assert.Equal(t, test.err, err)
			}
		})
	}
}

func saveDeploymentReturnsFail() Deployment {
	k8smock := repo_test.K8SMock{
		CreateDeploymentFunc: func(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error {
			return errors.New(" DB save failure")
		},
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := Deployment{K8S: &k8smock, DBconn: SetDB()}
	return dep
}

func saveDeploymentReturnsSuccess() Deployment {
	k8smock := repo_test.K8SMock{
		CreateDeploymentFunc: func(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error {
			return nil
		},
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := Deployment{K8S: &k8smock, DBconn: SetDB()}
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
