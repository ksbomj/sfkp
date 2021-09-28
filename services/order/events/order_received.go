package events

import (
	"github.com/google/uuid"
	"time"
)

// OrderReceived represents an order that was received in the system and will be published to our messaging system
type OrderReceived struct {
	BaseEvent
	Body interface{}
}

// GetID returns the unique identifier of the event
func (or OrderReceived) GetID() uuid.UUID {
	return or.ID
}

// GetName returns the name of the event
func (or OrderReceived) GetName() string {
	return "OrderReceived"
}

// GetTimestamp returns the unique timestamp of the event
func (or OrderReceived) GetTimestamp() time.Time {
	return or.Timestamp
}

// GetBody returns the body content of the event
func (or OrderReceived) GetBody() interface{} {
	return or.Body
}
