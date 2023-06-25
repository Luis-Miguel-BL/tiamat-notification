package marshaller

import (
	"encoding/json"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
)

type DomainEventMarshaller struct {
	eventTypeRegistry map[string]domain.DomainEvent
}

func NewDomainEventMarshaller() *DomainEventMarshaller {
	return &DomainEventMarshaller{
		eventTypeRegistry: map[string]domain.DomainEvent{
			string(event.ActionTriggedEventType):         event.ActionTrigged{},
			string(event.CustomerCreatedEventType):       event.CustomerCreatedEvent{},
			string(event.CustomerEventOccurredEventType): event.CustomerEventOccurredEvent{},
			string(event.CustomerMatchedEventType):       event.CustomerMatched{},
			string(event.StepJourneyFailedEventType):     event.StepJourneyFailed{},
			string(event.StepJourneyScheduledEventType):  event.StepJourneyScheduled{},
			string(event.StepJourneySkippedEventType):    event.StepJourneySkipped{},
			string(event.StepJourneySuccessedEventType):  event.StepJourneySuccessed{},
		},
	}
}

func MarshalDomainEvent(event domain.DomainEvent) (eventJson string, err error) {
	eventData, err := json.Marshal(map[string]interface{}{
		"event_type":     event.EventType(),
		"aggregate_type": event.AggregateType(),
		"aggregate_id":   event.AggregateID(),
		"occurred_at":    event.OccurredAt(),
		"data":           event,
	})

	if err != nil {
		return eventJson, err
	}

	return string(eventData), nil
}
func UnmarshalDomainEvent[T domain.DomainEvent](data string) (domainEvent T, err error) {
	err = json.Unmarshal([]byte(data), &domainEvent)
	if err != nil {
		return domainEvent, err
	}

	return domainEvent, nil
}
