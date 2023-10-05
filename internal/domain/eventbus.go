package domain

import "context"

type EventBusPublisher interface {
	Publish(ctx context.Context, event DomainEvent) (err error)
	Subscribe(ctx context.Context, messages chan DomainEvent, eventTypes ...EventType) (err error)
}

type EventBus interface {
	EventBusPublisher
}
