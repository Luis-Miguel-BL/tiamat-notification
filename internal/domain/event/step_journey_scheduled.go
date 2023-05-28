package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var StepJourneyScheduledType = domain.EventType("step-journey-scheduled")

type StepJourneyScheduled struct {
	*domain.DomainEventBase
	CustomerID    string
	WorkspaceID   string
	CampaignID    string
	ActionID      string
	StepJourneyID string
	TriggeredAt   time.Time
}
