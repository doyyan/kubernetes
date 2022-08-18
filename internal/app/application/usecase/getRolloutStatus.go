package usecase

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

//GetRolloutStatus makes a call to k8s to retrieve the current status of a deployment rollout
func GetRolloutStatus(ctx context.Context, logger *logrus.Logger, args CreateDeploymentArgs) (string, bool, error) {
	deployment := domain.Deployment{
		NameSpace: args.Deployment.Namespace,
		Name:      args.Deployment.Name,
	}
	return args.DeploymentRepository.GetRolloutStatus(ctx, logger, deployment)
}
