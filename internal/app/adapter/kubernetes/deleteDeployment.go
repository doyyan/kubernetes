package kubernetes

import (
	"context"

	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kube "k8s.io/client-go/kubernetes"
)

//Delete calls the k8s APIs to delete a deployment from k8s
func (k Kube) Delete(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error {
	deletePolicy := metav1.DeletePropagationForeground
	deploymentsClient := clientset.AppsV1().Deployments(d.NameSpace)
	if err := deploymentsClient.Delete(context.TODO(), d.Name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return err
	}
	return nil
}
