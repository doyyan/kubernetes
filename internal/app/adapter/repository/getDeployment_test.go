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

func Test_GetDeployment(t *testing.T) {
	tests := map[string]struct {
		err error
		dep Deployment
	}{
		"pass getDeploymentSuccess": {
			err: nil,
			dep: getDeploymentReturnsSuccess(),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.dep.Get(context.Background(), logrus.New(), domain.Deployment{})
			if test.err != nil {
				assert.Equal(t, test.err, err)
			}
		})
	}
}

func getDeploymentReturnsFail() Deployment {
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

func getDeploymentReturnsSuccess() Deployment {
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
