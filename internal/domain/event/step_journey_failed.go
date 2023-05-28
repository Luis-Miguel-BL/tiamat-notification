package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var StepJourneyFailedTypeType = domain.EventType("step-journey-failed")

type StepJourneyFailedType struct {
	*domain.DomainEventBase
	CustomerID    string
	WorkspaceID   string
	CampaignID    string
	ActionID      string
	StepJourneyID string
	Description   string
	TriggeredAt   time.Time
}
