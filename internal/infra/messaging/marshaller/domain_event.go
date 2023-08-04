package marshaller

import (
	"encoding/json"
	"fmt"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
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

func UnmarshalDomainEvent[T domain.DomainEvent](data string) (domainEvent T, err error) {
	err = json.Unmarshal([]byte(data), &domainEvent)
	if err != nil {
		return domainEvent, err
	}

	return domainEvent, nil
}
