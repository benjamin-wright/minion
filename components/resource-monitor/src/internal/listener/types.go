package listener

import "ponglehub.co.uk/resource-monitor/internal/monitors"

// Listener creates and deletes monitor cronjobs
type Listener struct {
	m *monitors.Monitors
}

// New create a new listener instance
func New() (Listener, error) {
	m, err := monitors.New()
	if err != nil {
		return Listener{}, err
	}

	return Listener{m: m}, nil
}
