package api

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

type EventConsumer interface {
	Consume(ctx context.Context, eventType domain.EventType, eventStr string) (err error)
}
