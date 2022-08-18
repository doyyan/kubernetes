package usecase

import (
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/doyyan/kubernetes/internal/app/domain/domainrepo"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type CreateDeploymentArgs struct {
	Deployment           valueObject.Deployment
	DeploymentRepository domainrepo.IDeploymentRepo
}

func CreateDeployment(ctx context.Context, logger *logrus.Logger, args CreateDeploymentArgs) error {
	deployment := domain.Deployment{
		NameSpace:     args.Deployment.Namespace,
		Name:          args.Deployment.Name,
		Kind:          "deployment",
		ContainerName: args.Deployment.ContainerName,
		ContainerPort: args.Deployment.ContainerPort,
		Image:         args.Deployment.Image,
		Labels:        args.Deployment.Labels,
		Replicas:      args.Deployment.Replicas,
	}
	if err := args.DeploymentRepository.Create(ctx, logger, deployment); err != nil {
		return err
	}
	return nil
}
