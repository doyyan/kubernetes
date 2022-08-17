package repository

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

type IDeploymentRepo interface {
	Get(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (domain.Deployment, error)
	Create(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error
	List(ctx context.Context, logger *logrus.Logger) ([]domain.Deployment, error)
	Delete(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error
	GetRolloutStatus(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (string, bool, error)
}
