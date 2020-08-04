package client

import (
	"k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create creates a new cronjob
func (w *Wrapper) Create(namespace string, cronjob *v1beta1.CronJob) (*v1beta1.CronJob, error) {
	return w.clientset.BatchV1beta1().CronJobs(namespace).Create(cronjob)
}

// Delete deletes an existing cronjob
func (w *Wrapper) Delete(namespace string, name string, options *metav1.DeleteOptions) error {
	return w.clientset.BatchV1beta1().CronJobs(namespace).Delete(name, options)
}
