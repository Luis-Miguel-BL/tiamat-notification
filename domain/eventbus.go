package domain

import "context"

type EventBusPublisher interface {
	Publish(ctx context.Context, event DomainEvent)
}

type EventBus interface {
	EventBusPublisher
}
