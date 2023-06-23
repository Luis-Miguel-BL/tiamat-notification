package domain

import "context"

type EventBusPublisher interface {
	Publish(ctx context.Context, event DomainEvent) (err error)
}

type EventBus interface {
	EventBusPublisher
}
