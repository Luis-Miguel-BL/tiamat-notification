package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var StepJourneyFailedEventType = domain.EventType("step-journey-failed")

type StepJourneyFailed struct {
	*domain.DomainEventBase
	CustomerID    string
	WorkspaceID   string
	CampaignID    string
	ActionID      string
	StepJourneyID string
	JourneyID     string
	Description   string
	TriggeredAt   time.Time
}
