package main

import (
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crdClient "ponglehub.co.uk/crd-lib/pkg/client"
	crdInformer "ponglehub.co.uk/crd-lib/pkg/informer"
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

	informer, err := crdInformer.Pipelines()
	if err != nil {
		log.Fatalf("Failed to create CRD informer: %+v", err)
	}

	for event := range informer.Events {
		log.Infof("Received event: %s", event.Kind)

		pipelines, err := client.ListPipelines(metav1.ListOptions{})
		if err != nil {
			log.Errorf("Failed to list pipelines: %+v", err)
		} else {
			log.Infof("Pipelines: %d", len(pipelines.Items))
		}
	}
}
