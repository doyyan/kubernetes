package repository

import "github.com/doyyan/kubernetes/internal/app/domain"

type ICustomer interface {
	Get(name string, kind string) domain.Deployment
	Create(deployment domain.Deployment)
	List() []domain.Deployment
	Delete(deployment domain.Deployment)
	GetStatus() string
}
