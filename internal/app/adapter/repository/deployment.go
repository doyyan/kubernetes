package repository

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	kube "k8s.io/client-go/kubernetes"
)

type K8S interface {
	GetKubeConfig() *kube.Clientset
	CreateDeployment(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error
}

type Deployment struct {
	DBconn *gorm.DB
	K8S    K8S
}

func (d Deployment) Create(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
	clientset := d.K8S.GetKubeConfig()
	err := d.K8S.CreateDeployment(ctx, logger, deployment, clientset)
	if err != nil {
		logger.Error(err)
		return err
	}
	val := model.JSONMap(deployment.Labels)
	dep := model.Deployment{
		Name:          deployment.Name,
		NameSpace:     deployment.Namespace,
		Kind:          deployment.Kind,
		Image:         deployment.Image,
		ContainerPort: deployment.ContainerPort,
		ContainerName: deployment.ContainerName,
		Replicas:      deployment.Replicas,
		LabelsDB:      val,
	}
	result := d.DBconn.Save(&dep)
	if result.Error != nil && result.RowsAffected != 1 {
		return result.Error
	}
	return nil
}
func (d Deployment) Get(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (domain.Deployment, error) {
	return domain.Deployment{}, nil
}
func (d Deployment) List(ctx context.Context, logger *logrus.Logger) ([]domain.Deployment, error) {
	return nil, nil
}
func (d Deployment) Delete(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
	return nil
}
func (d Deployment) GetStatus(ctx context.Context, logger *logrus.Logger) (domain.Deployment, error) {
	return domain.Deployment{}, nil
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
