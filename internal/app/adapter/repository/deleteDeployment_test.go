package repository

import (
	"errors"
	"testing"

	repo_test "github.com/doyyan/kubernetes/internal/app/adapter/repository/mocks"
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	kube "k8s.io/client-go/kubernetes"
)

func Test_DeleteDeployment(t *testing.T) {
	tests := map[string]struct {
		err error
		dep Deployment
	}{
		"pass deleteDeploymentSuccess": {
			err: nil,
			dep: deleteDeploymentReturnsSuccess(),
		},
		"pass deleteDeploymentFail": {
			err: errors.New(" DB save failure"),
			dep: deleteDeploymentReturnsFail(),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.dep.Delete(context.Background(), logrus.New(), domain.Deployment{})
			if test.err != nil {
				assert.Equal(t, test.err, err)
			}
		})
	}
}

func deleteDeploymentReturnsFail() Deployment {
	k8smock := repo_test.K8SMock{
		DeleteFunc: func(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error {
			return errors.New(" DB save failure")
		},
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := Deployment{K8S: &k8smock, DBconn: SetDB()}
	return dep
}

func deleteDeploymentReturnsSuccess() Deployment {
	k8smock := repo_test.K8SMock{
		DeleteFunc: func(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error {
			return nil
		},
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := Deployment{K8S: &k8smock, DBconn: SetDB()}
	return dep
}
