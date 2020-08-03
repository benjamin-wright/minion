package main

import (
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crdClient "ponglehub.co.uk/crd-lib/pkg/client"
	crdInformer "ponglehub.co.uk/crd-lib/pkg/informer"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
	"ponglehub.co.uk/resource-monitor/internal/config"
)

func main() {
	cfg := config.Get()
	log.Infof("Running with config: %s", cfg.String())

	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err == nil {
		log.SetLevel(logLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	err = v1alpha1.Init()
	if err != nil {
		log.Fatalf("Failed registering CRD with kube client: %+v", err)
	}

	client, err := crdClient.New()
	if err != nil {
		log.Fatalf("Failed to create CRD client: %+v", err)
	}

	informer, err := crdInformer.Resources()
	if err != nil {
		log.Fatalf("Failed to create resource CRD informer: %+v", err)
	}

	for event := range informer.Events {
		log.Infof("Received resource event: %s", event.Kind)

		resources, err := client.ListResources(metav1.ListOptions{})
		if err != nil {
			log.Errorf("Failed to list resources: %+v", err)
		} else {
			log.Infof("Resources: %d", len(resources.Items))
		}
	}
}
