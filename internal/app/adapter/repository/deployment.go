package repository

import (
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql"
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/domain"
)

type Deployment struct{}

func (d Deployment) Create(deployment domain.Deployment) error {
	db := postgresql.GetDB()
	dep := model.Deployment{
		Name:      deployment.Name,
		NameSpace: deployment.Namespace,
		Labels:    deployment.Labels,
		Replicas:  0,
	}
	dep.FillDefaults()
	result := db.Save(&dep)
	if result.Error != nil && result.RowsAffected != 1 {
		return result.Error
	}
	return nil
}
func (d Deployment) Get(domain.Deployment) (domain.Deployment, error) {
	return domain.Deployment{}, nil
}
func (d Deployment) List() ([]domain.Deployment, error) {
	return nil, nil
}
func (d Deployment) Delete(deployment domain.Deployment) error {
	return nil
}
func (d Deployment) GetStatus() (domain.Deployment, error) {
	return domain.Deployment{}, nil
}
