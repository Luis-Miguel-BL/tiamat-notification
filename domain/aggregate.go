package domain

import (
	"sync"
)

type AggregateID string
type AggregateType string

type Aggregate struct {
	AggregateID       AggregateID
	AggregateType     AggregateType
	UncommittedEvents []DomainEvent
	mu                *sync.Mutex
}

func (a *Aggregate) AppendEvent(event DomainEvent) {
	a.UncommittedEvents = append(a.UncommittedEvents, event)
}

func (a *Aggregate) CatchUncommitedEvents() []DomainEvent {
	a.mu.Lock()
	defer a.mu.Unlock()
	events := a.UncommittedEvents
	a.UncommittedEvents = nil
	return events
}
