package domain

import (
	"sync"
)

type AggregateID string
type AggregateType string

type AggregateRoot struct {
	aggregateID       AggregateID
	aggregateType     AggregateType
	uncommittedEvents []DomainEvent
	mu                *sync.Mutex
}

func NewAggregateRoot(aggregateType AggregateType, aggregateID AggregateID) *AggregateRoot {
	return &AggregateRoot{
		aggregateID:   aggregateID,
		aggregateType: aggregateType,
		mu:            &sync.Mutex{},
	}
}
func (a *AggregateRoot) AggregateID() AggregateID {
	return a.aggregateID
}

func (a *AggregateRoot) AggregateType() AggregateType {
	return a.aggregateType
}

func (a *AggregateRoot) AppendEvent(event DomainEvent) {
	a.uncommittedEvents = append(a.uncommittedEvents, event)
}

func (a *AggregateRoot) CatchUncommitedEvents() []DomainEvent {
	a.mu.Lock()
	defer a.mu.Unlock()
	events := a.uncommittedEvents
	a.uncommittedEvents = nil
	return events
}
