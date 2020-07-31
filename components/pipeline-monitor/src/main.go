package main

import (
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crdClient "ponglehub.co.uk/crd-lib/pkg/client"
	"ponglehub.co.uk/crd-lib/pkg/informer"
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

	pipelineInformer, err := crdInformer.Pipelines()
	if err != nil {
		log.Fatalf("Failed to create pipeline CRD informer: %+v", err)
	}

	resourceInformer, err := crdInformer.Resources()
	if err != nil {
		log.Fatalf("Failed to create resource CRD informer: %+v", err)
	}

	aggregated := make(chan interface{}, 20)

	go func(aggregated chan<- interface{}, pipelines <-chan informer.PipelineEvent) {
		for event := range pipelines {
			aggregated <- event
		}
	}(aggregated, pipelineInformer.Events)

	go func(aggregated chan<- interface{}, resources <-chan informer.ResourceEvent) {
		for event := range resources {
			aggregated <- event
		}
	}(aggregated, resourceInformer.Events)

	for event := range aggregated {
		switch event := event.(type) {
		case crdInformer.PipelineEvent:
			log.Infof("Received pipeline event: %s", event.Kind)

			pipelines, err := client.ListPipelines(metav1.ListOptions{})
			if err != nil {
				log.Errorf("Failed to list pipelines: %+v", err)
			} else {
				log.Infof("Pipelines: %d", len(pipelines.Items))
			}
		case crdInformer.ResourceEvent:
			log.Infof("Received resource event: %s", event.Kind)

			resources, err := client.ListResources(metav1.ListOptions{})
			if err != nil {
				log.Errorf("Failed to list resources: %+v", err)
			} else {
				log.Infof("Resources: %d", len(resources.Items))
			}
		default:
			log.Errorf("Unknown event type: %T", event)
		}
	}
}
