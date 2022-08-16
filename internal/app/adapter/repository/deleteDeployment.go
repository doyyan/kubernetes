package repository

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

func (d Deployment) Delete(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
	clientset := d.K8S.GetKubeConfig()
	err := d.K8S.Delete(ctx, logger, deployment, clientset)
	if err != nil {
		logger.Error(err)
		return err
	}
	dep := model.Deployment{
		Name:      deployment.Name,
		NameSpace: deployment.Namespace,
	}
	if err := d.DBconn.Where("name = ? AND namespace = ?", deployment.Name, deployment.Namespace).Delete(&dep).Error; err != nil {
		return err
	}
	return nil
}
