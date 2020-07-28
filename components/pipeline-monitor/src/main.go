package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crdClient "ponglehub.co.uk/crd-lib/pkg/client"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
	"ponglehub.co.uk/pipeline-monitor/config"
)

func main() {
	cfg := config.Get()
	log.Infof("Running with config: %s", cfg.String())

	err := v1alpha1.Init()
	if err != nil {
		log.Fatalf("Failed registering CRD with kube client: %+v", err)
	}

	client, err := crdClient.New()
	if err != nil {
		log.Fatalf("Failed to create CRD client: %+v", err)
	}

	for {
		pipelines, err := client.ListPipelines(metav1.ListOptions{})
		if err != nil {
			log.Errorf("Failed to list pipelines: %+v", err)
		} else {
			log.Infof("Pipelines: %+v", pipelines.Items)
		}

		time.Sleep(10 * time.Second)
	}
}
