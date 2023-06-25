package consumers

import "context"

type EventConsumer interface {
	Consume(ctx context.Context, eventType string, eventStr string) (err error)
}
