package repository

import (
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func (d Deployment) Get(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (domain.Deployment, error) {
	dep := model.Deployment{}
	d.DBconn.Where("name = ? AND namespace = ?", deployment.Name, deployment.NameSpace).Last(&dep)
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
	return out, nil
}
