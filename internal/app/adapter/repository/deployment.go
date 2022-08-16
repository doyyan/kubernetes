package repository

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/doyyan/kubernetes/internal/app/adapter/kubernetes"
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql"
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Deployment struct{}

func (d Deployment) Create(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
	clientset := kubernetes.GetKubeConfig()
	err := kubernetes.CreateDeployment(ctx, logger, deployment, clientset)
	if err != nil {
		logger.Error(err)
		return err
	}
	db := postgresql.GetDB()
	val := model.JSONMap(deployment.Labels)
	dep := model.Deployment{
		Name:          deployment.Name,
		NameSpace:     deployment.Namespace,
		Image:         deployment.Image,
		ContainerPort: deployment.ContainerPort,
		ContainerName: deployment.ContainerName,
		Replicas:      deployment.Replicas,
		LabelsDB:      val,
	}
	result := db.Save(&dep)
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
