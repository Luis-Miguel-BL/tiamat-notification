package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var ActionTriggedType = domain.EventType("action-trigged")

type ActionTrigged struct {
	*domain.DomainEventBase
	CustomerID        string
	WorkspaceID       string
	CampaignID        string
	ActionID          string
	CustomerJourneyID string
	TriggeredAt       time.Time
}
