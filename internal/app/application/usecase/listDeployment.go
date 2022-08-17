package usecase

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

func ListDeployment(ctx context.Context, logger *logrus.Logger, args CreateDeploymentArgs) ([]domain.Deployment, error) {
	dep, err := args.DeploymentRepository.List(ctx, logger)
	if err != nil {
		return nil, err
	}
	return dep, nil
}
