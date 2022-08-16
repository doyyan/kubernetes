package repository

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

func (d Deployment) GetStatus(ctx context.Context, logger *logrus.Logger) (domain.Deployment, error) {
	return domain.Deployment{}, nil
}
