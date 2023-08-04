package api

import "context"

type EventConsumer interface {
	Consume(ctx context.Context, eventStr string) (err error)
}
