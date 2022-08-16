package usecase

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

func DeleteDeployment(ctx context.Context, logger *logrus.Logger, args CreateDeploymentArgs) error {
	deployment := domain.Deployment{
		Namespace: args.Deployment.Namespace,
		Name:      args.Deployment.Name,
	}
	if err := args.DeploymentRepository.Delete(ctx, logger, deployment); err != nil {
		return err
	}
	return nil
}
