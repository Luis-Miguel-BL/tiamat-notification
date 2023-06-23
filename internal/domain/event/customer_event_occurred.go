package event

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var CustomerEventOccurredEventType = domain.EventType("customer-event-occurred")

type CustomerEventOccurredEvent struct {
	*domain.DomainEventBase
	CustomerID      string `json:"customer_id"`
	WorkspaceID     string `json:"workspace_id"`
	CustomerEventID string `json:"customer_event_id"`
}
