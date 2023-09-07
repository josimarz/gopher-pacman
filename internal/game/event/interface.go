package event

import "time"

type Event interface {
	GetName() string
	GetTimestamp() time.Time
	GetPayload() any
}

type EventHandler func(Event)
