package monitors

import (
	"ponglehub.co.uk/resource-monitor/internal/monitors/client"
)

// Monitors aggregation of monitoring related functions
type Monitors struct {
	client client.Interface
}

// New create a new version monitor
func New() (*Monitors, error) {
	client, err := client.New()
	if err != nil {
		return nil, err
	}

	return &Monitors{
		client: client,
	}, nil
}
