package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var CustomerUpdatedEventType = domain.EventType("customer-updated")

type CustomerUpdatedEvent struct {
	*domain.DomainEventBase
	CustomerID  string    `json:"customer_id"`
	WorkspaceID string    `json:"workspace_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
