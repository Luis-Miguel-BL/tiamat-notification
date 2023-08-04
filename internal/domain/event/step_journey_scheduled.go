package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var StepJourneyScheduledEventType = domain.EventType("step-journey-scheduled")

type StepJourneyScheduled struct {
	*domain.DomainEventBase
	CustomerID    string          `json:"customer_id"`
	WorkspaceID   string          `json:"workspace_id"`
	CampaignID    string          `json:"campaign_id"`
	ActionID      string          `json:"action_id"`
	StepJourneyID string          `json:"step_journey_id"`
	JourneyID     string          `json:"journey_id"`
	Reason        ScheduledReason `json:"reason"`
	TriggeredAt   time.Time       `json:"triggered_at"`
}

type ScheduledReason string

const (
	ScheduledReasonScheduledByAction          ScheduledReason = "scheduled-by-action"
	ScheduledReasonOutOfNotificationTimeRange ScheduledReason = "out-of-notification-time-range"
)
