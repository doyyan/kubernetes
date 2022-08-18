package kubernetes

import "github.com/doyyan/kubernetes/internal/app/domain"

// IDeploymentK8S an inteface for k8s kubectl processing
type IDeploymentK8S interface {
	Get(domain.Deployment) (domain.Deployment, error)
	Create(deployment domain.Deployment) error
	List() ([]domain.Deployment, error)
	Delete(deployment domain.Deployment) error
	GetStatus(domain.Deployment) (string, bool, error)
}
