package main

import (
	log "github.com/sirupsen/logrus"
	crdInformer "ponglehub.co.uk/crd-lib/pkg/informer"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
	"ponglehub.co.uk/resource-monitor/internal/config"
	"ponglehub.co.uk/resource-monitor/internal/listener"
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

	informer, err := crdInformer.Resources()
	if err != nil {
		log.Fatalf("Failed to create resource CRD informer: %+v", err)
	}

	listener, err := listener.New()
	if err != nil {
		log.Fatalf("Failed to create listener instance: %+v", err)
	}

	listener.Listen(informer.Events, cfg)
}
