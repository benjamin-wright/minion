package monitors

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
	"ponglehub.co.uk/resource-monitor/internal/config"
)

// Create create a new resource monitor
func (m *Monitors) Create(resource *v1alpha1.Resource, cfg config.Config) error {
	if resource == nil {
		return errors.New("Cannot create nil resource")
	}

	existing, err := m.client.Get(resource.Namespace, resource.Name+"-monitor", metav1.GetOptions{})
	if err != nil {
		return m.create(resource, cfg)
	}

	logrus.Infof("Monitor for '%s' already exists, checking for updates", resource.Name)
	return m.update(resource, existing, cfg)
}

// Update an existing resource
func (m *Monitors) Update(oldResource *v1alpha1.Resource, newResource *v1alpha1.Resource, cfg config.Config) error {
	if oldResource == nil {
		return errors.New("Cannot update from nil resource")
	}

	if newResource == nil {
		return errors.New("Cannot update to nil resource")
	}

	existing, err := m.client.Get(oldResource.Namespace, oldResource.Name+"-monitor", metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Failed to fetch existing monitor: %+v", err)
	}

	return m.update(newResource, existing, cfg)
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

func (m *Monitors) create(resource *v1alpha1.Resource, cfg config.Config) error {
	manifest, err := m.converter.Convert(resource, cfg)
	if err != nil {
		return fmt.Errorf("Failed to create monitor manifest: %+v", err)
	}

	_, err = m.client.Create(resource.ObjectMeta.Namespace, manifest)
	if err == nil {
		logrus.Infof("Monitor added for '%s'", resource.Name)
	}

	return err
}

func (m *Monitors) update(resource *v1alpha1.Resource, existing *v1beta1.CronJob, cfg config.Config) error {
	manifest, err := m.converter.Convert(resource, cfg)
	if err != nil {
		return fmt.Errorf("Failed to create monitor manifest: %+v", err)
	}

	existingSpec, err := m.converter.ConvertBack(existing)
	if err != nil {
		return fmt.Errorf("Failed to get resource spec from cronjob: %+v", err)
	}

	logrus.Infof("Existing spec: %+v", existingSpec)

	if existingSpec.Matches(resource.Spec) {
		logrus.Info("Skipping update, no changes")
		return nil
	}

	logrus.Infof("Changes found for %s", resource.Name)

	existing.ObjectMeta.Annotations = manifest.ObjectMeta.Annotations
	existing.ObjectMeta.Labels = manifest.ObjectMeta.Labels
	existing.Spec = manifest.Spec

	_, err = m.client.Update(resource.Namespace, existing)
	if err == nil {
		logrus.Infof("Monitor updated for '%s'", resource.Name)
	}
	return err
}
