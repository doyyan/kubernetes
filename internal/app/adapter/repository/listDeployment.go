package repository

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

func (d Deployment) List(ctx context.Context, logger *logrus.Logger) (deps []domain.Deployment, err error) {
	deployments := []model.Deployment{}
	result := d.DBconn.Find(&deployments)
	if result.Error != nil && result.RowsAffected != 1 {
		return nil, result.Error
	}
	for _, dep := range deployments {
		val := model.JSONMap(dep.LabelsDB)
		out := domain.Deployment{
			Name:          dep.Name,
			NameSpace:     dep.NameSpace,
			Kind:          dep.Kind,
			Image:         dep.Image,
			ContainerPort: dep.ContainerPort,
			ContainerName: dep.ContainerName,
			Replicas:      dep.Replicas,
			Labels:        val,
		}
		deps = append(deps, out)
	}
	return deps, nil
}
