package monitors

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
	"ponglehub.co.uk/resource-monitor/internal/config"
)

// Create create a new resource monitor
func (m *Monitors) Create(resource *v1alpha1.Resource, cfg config.Config) error {
	env := []corev1.EnvVar{}
	for _, variable := range resource.Spec.Env {
		env = append(env, corev1.EnvVar{
			Name:  variable.Name,
			Value: variable.Value,
		})
	}

	volumes := []corev1.Volume{}
	mounts := []corev1.VolumeMount{}
	for _, secret := range resource.Spec.Secrets {
		items := []corev1.KeyToPath{}
		for _, key := range secret.Keys {
			items = append(items, corev1.KeyToPath{
				Key:  key.Key,
				Path: key.Path,
			})
		}

		volumes = append(volumes, corev1.Volume{
			Name: secret.Name,
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: secret.Name,
					Items:      items,
				},
			},
		})

		mounts = append(mounts, corev1.VolumeMount{
			Name:      secret.Name,
			MountPath: "",
		})
	}

	volumes = append(volumes, corev1.Volume{
		Name: "results",
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{
				Medium: corev1.StorageMediumMemory,
			},
		},
	})

	_, err := m.cronjobs.Create(&batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-monitor", resource.ObjectMeta.Name),
			Namespace: resource.ObjectMeta.Namespace,
		},
		Spec: batchv1beta1.CronJobSpec{
			JobTemplate: batchv1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{
								{
									Name:         "check",
									Image:        resource.Spec.Image,
									Env:          env,
									VolumeMounts: mounts,
								},
							},
							Containers: []corev1.Container{
								{
									Name:  "sidecar",
									Image: cfg.SidecarImage,
									Env: []corev1.EnvVar{
										{
											Name:  "RESOURCE",
											Value: resource.ObjectMeta.Name,
										},
										{
											Name:  "LOG_LEVEL",
											Value: cfg.LogLevel,
										},
									},
								},
							},
							Volumes: volumes,
						},
					},
				},
			},
		},
	})

	return err
}
