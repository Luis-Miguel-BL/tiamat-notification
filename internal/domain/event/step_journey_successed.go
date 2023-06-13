package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var StepJourneySuccessedEventType = domain.EventType("step-journey-successed")

type StepJourneySuccessed struct {
	*domain.DomainEventBase
	CustomerID    string
	WorkspaceID   string
	CampaignID    string
	ActionID      string
	StepJourneyID string
	JourneyID     string
	TriggeredAt   time.Time
}
