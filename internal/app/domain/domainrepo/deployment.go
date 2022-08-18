package domainrepo

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

//go:generate moq -out mocks/deploymentRepo.mock.go -pkg domainrepotest -skip-ensure . IDeploymentRepo
// IDeploymentRepo an interface which will be dependency Injected by the adapter to get the
// usecase to call methods on the Deployment domain
type IDeploymentRepo interface {
	Get(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (domain.Deployment, error)
	Create(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error
	List(ctx context.Context, logger *logrus.Logger) ([]domain.Deployment, error)
	Delete(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error
	GetRolloutStatus(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) (string, bool, error)
}
