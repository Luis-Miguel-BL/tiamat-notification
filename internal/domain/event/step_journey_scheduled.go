package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var StepJourneyScheduledEventType = domain.EventType("step-journey-scheduled")

type StepJourneyScheduled struct {
	*domain.DomainEventBase
	CustomerID    string
	WorkspaceID   string
	CampaignID    string
	ActionID      string
	StepJourneyID string
	JourneyID     string
	Reason        ScheduledReason
	TriggeredAt   time.Time
}

type ScheduledReason string

const (
	ScheduledReasonScheduledByAction          ScheduledReason = "scheduled-by-action"
	ScheduledReasonOutOfNotificationTimeRange ScheduledReason = "out-of-notification-time-range"
)
