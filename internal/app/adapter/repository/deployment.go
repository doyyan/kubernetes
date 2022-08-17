package repository

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	kube "k8s.io/client-go/kubernetes"
)

//go:generate moq -out mocks/k8s.mock.go -pkg repo_test -skip-ensure . K8S
type K8S interface {
	GetKubeConfig() *kube.Clientset
	CreateDeployment(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error
	Delete(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment, clientset *kube.Clientset) error
	GetRolloutStatus(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment, clientset *kube.Clientset) (string, bool, error)
}
type Deployment struct {
	DBconn *gorm.DB
	K8S    K8S
}

// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func Value(labels map[string]string) (driver.Value, error) {
	return json.Marshal(labels)
}
func (d Deployment) setDB(dbconn *gorm.DB) {
	d.DBconn = dbconn
}
func (d Deployment) setKube(kube K8S) {
	d.K8S = kube
}
