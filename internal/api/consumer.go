package api

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

type EventConsumer interface {
	Consume(ctx context.Context, event domain.DomainEvent) (err error)
}
