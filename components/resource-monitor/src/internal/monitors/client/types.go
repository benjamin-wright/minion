package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Wrapper wraps around the kube clientset
type Wrapper struct {
	clientset *kubernetes.Clientset
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
