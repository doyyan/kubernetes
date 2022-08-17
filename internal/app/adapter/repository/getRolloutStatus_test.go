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

func Test_RolloutStatus(t *testing.T) {
	tests := map[string]struct {
		err     error
		dep     Deployment
		message string
		status  bool
	}{
		"pass getRolloutStatusSuccess": {
			err:     nil,
			dep:     getRolloutStatusReturnsSuccess(),
			message: "success",
			status:  true,
		},
		"pass getRolloutStagtusFail": {
			err:     errors.New(" DB save failure"),
			dep:     getRolloutStatusReturnsFail(),
			message: "",
			status:  false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			message, status, err := test.dep.GetRolloutStatus(context.Background(), logrus.New(), domain.Deployment{})
			if test.err != nil {
				assert.Equal(t, test.err, err)
				assert.Equal(t, test.status, status)
				assert.Equal(t, test.message, message)
			}
		})
	}
}

func getRolloutStatusReturnsFail() Deployment {
	k8smock := repo_test.K8SMock{
		GetRolloutStatusFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment, clientset *kube.Clientset) (string, bool, error) {
			return "", false, errors.New(" DB save failure")
		},
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := Deployment{K8S: &k8smock, DBconn: SetDB()}
	return dep
}

func getRolloutStatusReturnsSuccess() Deployment {
	k8smock := repo_test.K8SMock{
		GetRolloutStatusFunc: func(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment, clientset *kube.Clientset) (string, bool, error) {
			return "success", true, nil
		},
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := Deployment{K8S: &k8smock, DBconn: SetDB()}
	return dep
}
