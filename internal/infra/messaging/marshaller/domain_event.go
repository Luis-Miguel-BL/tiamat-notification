package marshaller

import (
	"encoding/json"
	"fmt"

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

func (m *DomainEventMarshaller) Marshal(event domain.DomainEvent) (eventJson string, err error) {
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
func (m *DomainEventMarshaller) Unmarshal(data string) (domainEvent domain.DomainEvent, err error) {
	rawEvent := map[string]any{}
	err = json.Unmarshal([]byte(data), &rawEvent)

	if err != nil {
		return domainEvent, err
	}

	eventType, ok := rawEvent["event_type"].(string)
	if !ok {
		return domainEvent, fmt.Errorf("invalid event_type field: %s", string(data))
	}

	domainEvent, ok = m.eventTypeRegistry[eventType]
	if !ok {
		return domainEvent, fmt.Errorf("invalid event_type field %s", string(data))
	}

	eventDataJson, _ := json.Marshal(rawEvent["data"])

	err = json.Unmarshal(eventDataJson, domainEvent)
	if !ok {
		return domainEvent, err
	}

	return domainEvent, nil
}
