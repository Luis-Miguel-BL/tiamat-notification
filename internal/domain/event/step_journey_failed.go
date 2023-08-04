package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var StepJourneyFailedEventType = domain.EventType("step-journey-failed")

type StepJourneyFailed struct {
	*domain.DomainEventBase
	CustomerID    string    `json:"customer_id"`
	WorkspaceID   string    `json:"workspace_id"`
	CampaignID    string    `json:"campaign_id"`
	ActionID      string    `json:"action_id"`
	StepJourneyID string    `json:"step_journey_id"`
	JourneyID     string    `json:"journey_id"`
	Description   string    `json:"description"`
	TriggeredAt   time.Time `json:"triggered_at"`
}
