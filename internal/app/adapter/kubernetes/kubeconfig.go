package kubernetes

import (
	"flag"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	kube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var clientset *kube.Clientset

//SetConfig initialises connection to a k8s API server
func (k Kube) SetConfig(ctx context.Context, logger *logrus.Logger) error {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logger.Error(err)
	}
	clientset, err = kube.NewForConfig(config)
	return err
}

//GetKubeConfig returns the k8s config
func (k Kube) GetKubeConfig() *kube.Clientset {
	return clientset
}

//Int32Ptr converts an int32 to a int32 pointer
func Int32Ptr(i int32) *int32 { return &i }
