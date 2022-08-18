package usecase

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

//GetDeployment calls the DB to send query data to retrieve a single deployment
func GetDeployment(ctx context.Context, logger *logrus.Logger, args CreateDeploymentArgs) (domain.Deployment, error) {
	deployment := domain.Deployment{
		NameSpace: args.Deployment.Namespace,
		Name:      args.Deployment.Name,
	}
	dep, err := args.DeploymentRepository.Get(ctx, logger, deployment)
	if err != nil {
		return domain.Deployment{}, err
	}
	return dep, nil
}
