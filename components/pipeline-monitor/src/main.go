package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crdClient "ponglehub.co.uk/crd-lib/pkg/client"
	"ponglehub.co.uk/pipeline-monitor/config"
)

func main() {
	cfg := config.Get()
	log.Infof("Running with config: %s", cfg.String())

	client, err := crdClient.New()
	if err != nil {
		log.Fatalf("Failed to create CRD client: %+v", err)
	}

	for {
		pipelines, err := client.ListPipelines(metav1.ListOptions{})
		if err != nil {
			log.Errorf("Failed to list pipelines: %+v", err)
		} else {
			log.Infof("Pipelines: %+v", pipelines)
		}

		time.Sleep(10 * time.Second)
	}
}
