package monitors

import (
	"k8s.io/client-go/kubernetes"
	crons "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	"k8s.io/client-go/rest"
)

// Monitors aggregation of monitoring related functions
type Monitors struct {
	cronjobs crons.CronJobInterface
}

// New create a new version monitor
func New(namespace string) (*Monitors, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Monitors{
		cronjobs: clientset.BatchV1beta1().CronJobs(namespace),
	}, nil
}
