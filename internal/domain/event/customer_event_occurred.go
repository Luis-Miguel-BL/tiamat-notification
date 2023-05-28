package event

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

var CustomerEventOccurredEventType = domain.EventType("customer-event-occurred")

type CustomerEventOccurredEvent struct {
	*domain.DomainEventBase
	CustomerID       string
	WorkspaceID      string
	CustomerEventID  string
	ExternalID       vo.ExternalID
	Slug             vo.Slug
	CustomAttributes vo.CustomAttributes
}
