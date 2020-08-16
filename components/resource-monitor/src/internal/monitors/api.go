package monitors

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
	"ponglehub.co.uk/resource-monitor/internal/config"
)

// Update an existing resource
func (m *Monitors) Update(oldResource *v1alpha1.Resource, newResource *v1alpha1.Resource, cfg config.Config) error {
	if oldResource == nil {
		return errors.New("Cannot update from nil resource")
	}

	if newResource == nil {
		return errors.New("Cannot update to nil resource")
	}

	existing, err := m.client.Get(oldResource.Namespace, oldResource.Name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Failed to fetch existing monitor: %+v", err)
	}

	existingSpec, err := m.converter.ConvertBack(existing)
	if err != nil {
		return fmt.Errorf("Failed to get resource spec from cronjob: %+v", err)
	}

	if existingSpec.Matches(newResource.Spec) {
		logrus.Info("Skipping update, no changes")
		return nil
	}

	manifest, err := m.converter.Convert(newResource, cfg)
	if err != nil {
		return fmt.Errorf("Failed to get manifest for new resource monitor: %+v", err)
	}

	existing.ObjectMeta.Annotations = manifest.ObjectMeta.Annotations
	existing.ObjectMeta.Labels = manifest.ObjectMeta.Labels
	existing.Spec = manifest.Spec

	_, err = m.client.Update(newResource.Namespace, existing)
	return err
}

// Delete deletes the monitor for a given resource
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

	manifest, err := m.converter.Convert(resource, cfg)
	if err != nil {
		return fmt.Errorf("Failed to create new resource monitor: %+v", err)
	}

	_, err = m.client.Create(resource.ObjectMeta.Namespace, manifest)
	return err
}
