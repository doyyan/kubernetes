package repository

import (
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

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
		NameSpace:     deployment.NameSpace,
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
