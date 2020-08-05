package monitors

import (
	"k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
	"ponglehub.co.uk/resource-monitor/internal/config"
	"ponglehub.co.uk/resource-monitor/internal/monitors/client"
	"ponglehub.co.uk/resource-monitor/internal/monitors/converter"
)

// Interface behaviour contract for the cronjobs api
type clientInterface interface {
	Create(namespace string, cronjob *v1beta1.CronJob) (*v1beta1.CronJob, error)
	Delete(namespace string, name string, options *metav1.DeleteOptions) error
	Get(namespace string, name string, options metav1.GetOptions) (*v1beta1.CronJob, error)
}

type converterInterface interface {
	Convert(resource *v1alpha1.Resource, cfg config.Config) *v1beta1.CronJob
}

// Monitors aggregation of monitoring related functions
type Monitors struct {
	client    clientInterface
	converter converterInterface
}

// New create a new version monitor
func New() (*Monitors, error) {
	client, err := client.New()
	if err != nil {
		return nil, err
	}

	return &Monitors{
		client:    client,
		converter: &converter.Converter{},
	}, nil
}
