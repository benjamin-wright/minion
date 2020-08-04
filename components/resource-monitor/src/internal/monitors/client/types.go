package client

import (
	"k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Wrapper wraps around the kube clientset
type Wrapper struct {
	clientset *kubernetes.Clientset
}

// Interface behaviour contract for the cronjobs api
type Interface interface {
	Create(namespace string, cronjob *v1beta1.CronJob) (*v1beta1.CronJob, error)
	Delete(namespace string, name string, options *metav1.DeleteOptions) error
}

// New create a new version monitor
func New() (*Wrapper, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Wrapper{
		clientset: clientset,
	}, nil
}
