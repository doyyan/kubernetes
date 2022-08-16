package kubernetes

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
)

func (k Kube) Delete(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment) error {
	//TODO implement me
	panic("implement me")
}
