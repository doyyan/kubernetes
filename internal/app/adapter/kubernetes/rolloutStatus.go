package kubernetes

import (
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	kube "k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/polymorphichelpers"
)

func (k Kube) GetRolloutStatus(ctx context.Context, logger *logrus.Logger, deployment domain.Deployment, clientset *kube.Clientset) (string, bool, error) {
	deploymentsClient := clientset.AppsV1().Deployments(deployment.NameSpace)
	dep, err := deploymentsClient.Get(context.TODO(), deployment.Name, metav1.GetOptions{})
	if err != nil {
		return "", false, err
	}
	unstructuredD := &unstructured.Unstructured{}
	unstructuredD.Object, err = runtime.DefaultUnstructuredConverter.ToUnstructured(dep)
	if err != nil {
		logger.Error(err)
	}
	dsv := polymorphichelpers.DeploymentStatusViewer{}
	return dsv.Status(unstructuredD, 0)
}
