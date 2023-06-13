package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var ActionTriggedEventType = domain.EventType("action-trigged")

type ActionTrigged struct {
	*domain.DomainEventBase
	CustomerID    string
	WorkspaceID   string
	CampaignID    string
	ActionID      string
	StepJourneyID string
	TriggeredAt   time.Time
}
