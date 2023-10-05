package marshaller

import (
	"encoding/json"
	"fmt"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
)

func UnmarshalDomainEventType(eventStr string) (eventType domain.EventType, err error) {
	mapEvent := make(map[string]interface{})

	err = json.Unmarshal([]byte(eventStr), &mapEvent)
	if err != nil {
		return eventType, err
	}
	eventType, ok := mapEvent["eventType"].(domain.EventType)
	if !ok {
		return eventType, fmt.Errorf("eventType not found %s", eventStr)
	}

	return eventType, nil
}

func UnmarshalDomainEvent(eventType domain.EventType, data string) (domainEvent domain.DomainEvent, err error) {
	switch eventType {
	case event.ActionTriggedEventType:
		domainEvent, err = unmarshalDomainEvent[event.ActionTriggedEvent](data)
	case event.CustomerCreatedEventType:
		domainEvent, err = unmarshalDomainEvent[event.CustomerCreatedEvent](data)
	case event.CustomerEventOccurredEventType:
		domainEvent, err = unmarshalDomainEvent[event.CustomerEventOccurredEvent](data)
	case event.CustomerMatchedEventType:
		domainEvent, err = unmarshalDomainEvent[event.CustomerMatchedEvent](data)
	case event.CustomerUpdatedEventType:
		domainEvent, err = unmarshalDomainEvent[event.CustomerUpdatedEvent](data)
	case event.StepJourneyFailedEventType:
		domainEvent, err = unmarshalDomainEvent[event.StepJourneyFailed](data)
	case event.StepJourneyScheduledEventType:
		domainEvent, err = unmarshalDomainEvent[event.StepJourneyScheduled](data)
	case event.StepJourneySkippedEventType:
		domainEvent, err = unmarshalDomainEvent[event.StepJourneySkipped](data)
	case event.StepJourneySuccessedEventType:
		domainEvent, err = unmarshalDomainEvent[event.StepJourneySuccessed](data)
	}
	if err != nil {
		return domainEvent, err
	}

	return domainEvent, nil
}

func unmarshalDomainEvent[T domain.DomainEvent](data string) (domainEvent T, err error) {
	err = json.Unmarshal([]byte(data), &domainEvent)
	if err != nil {
		return domainEvent, err
	}
	return domainEvent, nil
}
