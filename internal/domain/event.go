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
	eventType     EventType
	occurredAt    time.Time
	aggregateType AggregateType
	aggregateID   AggregateID
}
type NewDomainEventBaseInput struct {
	EventType     EventType
	OccurredAt    time.Time
	AggregateType AggregateType
	AggregateID   AggregateID
}

func NewDomainEventBase(input NewDomainEventBaseInput) *DomainEventBase {
	return &DomainEventBase{
		eventType:     input.EventType,
		occurredAt:    input.OccurredAt,
		aggregateType: input.AggregateType,
		aggregateID:   input.AggregateID,
	}
}
func (e DomainEventBase) EventType() EventType {
	return e.eventType
}
func (e DomainEventBase) OccurredAt() time.Time {
	return e.occurredAt
}
func (e DomainEventBase) AggregateType() AggregateType {
	return e.aggregateType
}
func (e DomainEventBase) AggregateID() AggregateID {
	return e.aggregateID
}
