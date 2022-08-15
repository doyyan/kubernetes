package usecase

import (
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/doyyan/kubernetes/internal/app/domain/repository"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
)

type CreateDeploymentArgs struct {
	Deployment           valueObject.Deployment
	DeploymentRepository repository.IDeploymentRepo
}

func CreateDeployment(args CreateDeploymentArgs) error {
	deployment := domain.Deployment{
		Namespace: args.Deployment.Namespace,
		Name:      args.Deployment.Name,
		Labels:    args.Deployment.Labels,
		Replicas:  args.Deployment.Replicas,
	}
	if err := args.DeploymentRepository.Create(deployment); err != nil {
		return err
	}
	return nil
}
