package monitors

import (
	"fmt"

	v1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	crons "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	"k8s.io/client-go/rest"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
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

// Create create a new resource monitor
func (m *Monitors) Create(resource *v1alpha1.Resource) {
	m.cronjobs.Create(&v1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-monitor", resource.ObjectMeta.Name),
			Namespace: resource.ObjectMeta.Namespace,
		},
		Spec: v1beta1.CronJobSpec{
			JobTemplate: v1beta1.JobTemplateSpec{
				Spec: v1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{},
							Containers:     []corev1.Container{},
							Volumes: []corev1.Volume{
								{
									Name: "results",
									VolumeSource: corev1.VolumeSource{
										EmptyDir: &corev1.EmptyDirVolumeSource{
											Medium: corev1.StorageMediumMemory,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
}
