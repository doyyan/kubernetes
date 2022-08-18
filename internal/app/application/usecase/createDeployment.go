package usecase

import (
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/doyyan/kubernetes/internal/app/domain/domainrepo"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

//CreateDeploymentArgs arguments needed to transport a deployment from the domain/usecase
// layer to the addpter layer and an interface for deployment DB methods
type CreateDeploymentArgs struct {
	Deployment           valueObject.Deployment
	DeploymentRepository domainrepo.IDeploymentRepo
}

//CreateDeployment sends a call to k8s to create a deployment and on success insert deployment
// data to the database
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
