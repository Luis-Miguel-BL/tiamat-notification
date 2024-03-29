package messaging

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

type AggregateEventDispatcher struct {
	Eventbus domain.EventBus
}

func NewAggregateEventDispatcher(eventbus domain.EventBus) *AggregateEventDispatcher {
	return &AggregateEventDispatcher{
		Eventbus: eventbus,
	}
}

func (d *AggregateEventDispatcher) PublishUncommitedEvents(ctx context.Context, aggregate domain.AggregateRoot) (err error) {
	uncommittedEvents := aggregate.CatchUncommitedEvents()
	for _, event := range uncommittedEvents {
		err = d.Eventbus.Publish(ctx, event)
		if err != nil {
			return err
		}
	}
	return nil
}
