package kubernetes

import "github.com/doyyan/kubernetes/internal/app/domain"

type IDeploymentK8S interface {
	Get(domain.Deployment) (domain.Deployment, error)
	Create(deployment domain.Deployment) error
	List() ([]domain.Deployment, error)
	Delete(deployment domain.Deployment) error
	GetStatus() (domain.Deployment, error)
}
