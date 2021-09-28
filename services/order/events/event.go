package events

import (
	"github.com/google/uuid"
	"time"
)

// Event is the interface for events
type Event interface {
	GetID() uuid.UUID
	GetName() string
	GetTimestamp() time.Time
	GetBody() interface{}
}

// BaseEvent represents common properties of an event
type BaseEvent struct {
	ID        uuid.UUID
	Timestamp time.Time
}
