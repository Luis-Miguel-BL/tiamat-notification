package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

var CustomerCreatedEventType = domain.EventType("customer-created")

type CustomerCreatedEvent struct {
	*domain.DomainEventBase
	CustomerID       string
	WorkspaceID      string
	Name             vo.PersonName
	Contact          vo.Contact
	CustomAttributes vo.CustomAttributes
	CreatedAt        time.Time
}
