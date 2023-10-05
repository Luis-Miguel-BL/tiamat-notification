package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var ActionTriggedEventType = domain.EventType("action-trigged")

type ActionTriggedEvent struct {
	*domain.DomainEventBase
	CustomerID    string    `json:"customer_id"`
	WorkspaceID   string    `json:"workspace_id"`
	CampaignID    string    `json:"campaign_id"`
	ActionID      string    `json:"action_id"`
	StepJourneyID string    `json:"step_journey_id"`
	TriggeredAt   time.Time `json:"triggered_at"`
}
