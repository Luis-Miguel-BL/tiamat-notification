package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var CustomerCreatedEventType = domain.EventType("customer-created")

type CustomerCreatedEvent struct {
	*domain.DomainEventBase
	CustomerID  string    `json:"customer_id"`
	WorkspaceID string    `json:"workspace_id"`
	CreatedAt   time.Time `json:"created_at"`
}
