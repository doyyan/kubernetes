package repository

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

func (d Deployment) GetRolloutStatus(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (string, bool, error) {
	clientset := d.K8S.GetKubeConfig()
	return d.K8S.GetRolloutStatus(ctx, logger, deployment, clientset)
}
