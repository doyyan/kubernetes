package repository

import (
	"testing"

	repo_test "github.com/doyyan/kubernetes/internal/app/adapter/repository/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	kube "k8s.io/client-go/kubernetes"
)

func Test_ListDeploymentsStatus(t *testing.T) {
	tests := map[string]struct {
		err     error
		dep     Deployment
		message string
		status  bool
	}{
		"pass listDeploymentsSuccess": {
			err:     nil,
			dep:     listDeploymentReturnsSuccess(),
			message: "success",
			status:  true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.dep.List(context.Background(), logrus.New())
			if test.err != nil {
				assert.Equal(t, test.err, err)
			}
		})
	}
}

func listDeploymentReturnsSuccess() Deployment {
	k8smock := repo_test.K8SMock{
		GetKubeConfigFunc: func() *kube.Clientset {
			return &kube.Clientset{}
		},
	}
	dep := Deployment{K8S: &k8smock, DBconn: SetDB()}
	return dep
}
