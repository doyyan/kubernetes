package repository

import (
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	kube "k8s.io/client-go/kubernetes"
)

//go:generate moq -out mocks/k8s.mock.go -pkg repo_test -skip-ensure . K8S

//K8S main interface which is implemented to do the k8s deployment operations
type K8S interface {
	GetKubeConfig() *kube.Clientset
	CreateDeployment(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error
	Delete(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment, clientset *kube.Clientset) error
	GetRolloutStatus(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment, clientset *kube.Clientset) (string, bool, error)
}

//Deployment main object that stores the handles to the k8s and DB adapters
type Deployment struct {
	DBconn *gorm.DB
	K8S    K8S
}
