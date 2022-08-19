package kubernetes

import (
	"github.com/doyyan/kubernetes/internal/app/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kube "k8s.io/client-go/kubernetes"
)

type Kube struct{}

// CreateDeployment calls the k8s APIs to create a deployment
func (k Kube) CreateDeployment(ctx context.Context, logger *logrus.Logger, d domain.Deployment, clientset *kube.Clientset) error {
	deploymentsClient := clientset.AppsV1().Deployments(d.NameSpace)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: d.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: Int32Ptr(int32(d.Replicas)),
			Selector: &metav1.LabelSelector{
				MatchLabels: d.Labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: d.Labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  d.ContainerName,
							Image: d.Image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: int32(d.ContainerPort),
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("Created deployment %q.\n", result.GetObjectMeta().GetName())
	return nil
}
