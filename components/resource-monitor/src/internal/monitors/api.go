package monitors

import (
	"errors"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
	"ponglehub.co.uk/resource-monitor/internal/config"
)

// Update an existing resource
func (m *Monitors) Update(oldResource *v1alpha1.Resource, newResource *v1alpha1.Resource) error {
	return nil
}

// Delete delete the monitor for a given resource
func (m *Monitors) Delete(resource *v1alpha1.Resource) error {
	if resource == nil {
		return errors.New("Cannot delete nil resource")
	}

	err := m.client.Delete(
		resource.ObjectMeta.Namespace,
		fmt.Sprintf("%s-monitor", resource.ObjectMeta.Name),
		&metav1.DeleteOptions{},
	)

	return err
}

// Create create a new resource monitor
func (m *Monitors) Create(resource *v1alpha1.Resource, cfg config.Config) error {
	if resource == nil {
		return errors.New("Cannot create nil resource")
	}

	manifest := m.converter.Convert(resource, cfg)
	_, err := m.client.Create(resource.ObjectMeta.Namespace, manifest)

	return err
}
