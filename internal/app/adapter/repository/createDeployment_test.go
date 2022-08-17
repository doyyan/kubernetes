package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	repo_test "github.com/doyyan/kubernetes/internal/app/adapter/repository/mocks"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	kube "k8s.io/client-go/kubernetes"
)

type Student struct {
	//*gorm.Model
	Name string
	ID   string
}
type depSuite struct {
	db         *gorm.DB
	mock       sqlmock.Sqlmock
	deployment model.Deployment
}

func Test(t *testing.T) {
	tests := map[string]struct {
		err error
	}{
		"pass createDeploymentSuccess": {
			nil,
		},
	}
	for name, _ := range tests {
		t.Run(name, func(t *testing.T) {
			dep := saveDeploymentReturnsSuccess()
			dep.Create(context.Background(), logrus.New(), domain.Deployment{})
		})
	}
}

func saveDeploymentReturnsSuccess() Deployment {
	db1, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db1,
		PreferSimpleProtocol: true,
	})
	db2, _ := gorm.Open(dialector, &gorm.Config{})
	k8smock := repo_test.K8SMock{
		CreateDeploymentFunc: func(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error {
			return nil
		},
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := Deployment{K8S: &k8smock, DBconn: db2}
	return dep
}
