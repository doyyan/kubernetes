package usecase

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

//DeleteDeployment sends a k8s call to delete a deployment and on success deletes
// the deployment from the database
func DeleteDeployment(ctx context.Context, logger *logrus.Logger, args CreateDeploymentArgs) error {
	deployment := domain.Deployment{
		NameSpace: args.Deployment.Namespace,
		Name:      args.Deployment.Name,
	}
	if err := args.DeploymentRepository.Delete(ctx, logger, deployment); err != nil {
		return err
	}
	return nil
}
