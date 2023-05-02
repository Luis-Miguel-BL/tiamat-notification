package domain

import (
	"time"
)

type EventType string

type DomainEvent interface {
	EventType() EventType
	OccurredAt() time.Time
	AggregateType() AggregateType
	AggregateID() AggregateID
}

type DomainEventBase struct {
	EventType     EventType
	OccurredAt    time.Time
	AggregateType AggregateType
	AggregateID   AggregateID
}
