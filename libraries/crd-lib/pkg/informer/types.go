package informer

// EventType the kind of event
type EventType string

const (
	ADDED   EventType = "ADDED"
	UPDATED EventType = "UPDATED"
	DELETED EventType = "DELETED"
)
