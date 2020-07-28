package client

import (
	"k8s.io/client-go/rest"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
)

// MinionCRDClient a wrapper to simplify interacting with Minion-CI CRDs
type MinionCRDClient struct {
	client *rest.RESTClient
}

// New create a new client instance
func New() (MinionCRDClient, error) {
	client, err := v1alpha1.TypedClient()
	if err != nil {
		return MinionCRDClient{}, err
	}

	return MinionCRDClient{
		client: client,
	}, nil
}
