package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crds "ponglehub.co.uk/crd-lib/pkg/client"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
	"ponglehub.co.uk/version-sidecar/internal/config"
	"ponglehub.co.uk/version-sidecar/internal/loading"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("Failed to load environment: %+v", err)
	}

	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err == nil {
		log.SetLevel(logLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Info("Running...")
	err = v1alpha1.Init()
	if err != nil {
		log.Fatalf("Failed to initialise kube client: %+v", err)
	}

	log.Info("Got client")
	client, err := crds.New()
	if err != nil {
		log.Fatalf("Failed to create client: %+v", err)
	}

	version, err := loading.LoadVersion()
	if err != nil {
		log.Fatalf("Failed to create client: %+v", err)
	}

	err = client.PostVersion(&v1alpha1.Version{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-%s", cfg.Resource, version),
			Labels: map[string]string{
				"resource": cfg.Resource,
			},
		},
		Spec: v1alpha1.VersionSpec{
			Resource: cfg.Resource,
			Version:  version,
		},
	}, "default")

	if err != nil {
		log.Fatalf("Failed to create version: %+v", err)
	}
}
