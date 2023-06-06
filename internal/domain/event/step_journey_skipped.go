package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var StepJourneySkippedType = domain.EventType("step-journey-skipped")

type StepJourneySkipped struct {
	*domain.DomainEventBase
	CustomerID    string
	WorkspaceID   string
	CampaignID    string
	ActionID      string
	StepJourneyID string
	Reason        SkippedReason
	TriggeredAt   time.Time
}

type SkippedReason string

const (
	SkippedReasonMatchFilters   SkippedReason = "campaign-filter-matched"
	SkippedReasonActionDisabled SkippedReason = "action-disabled"
)
