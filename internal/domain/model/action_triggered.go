package model

import (
	"time"
)

type ActionTriggeredID string

func NewActionTriggeredID(actionTriggeredID string) ActionTriggeredID {
	return ActionTriggeredID(actionTriggeredID)
}

type ActionTriggeredStatus string

const (
	ActionTriggeredStatusSuccess   ActionTriggeredStatus = "success"
	ActionTriggeredStatusScheduled ActionTriggeredStatus = "scheduled"
	ActionTriggeredStatusFailed    ActionTriggeredStatus = "failed"
)

type ActionTriggered struct {
	actionTriggeredID ActionTriggeredID
	campaignID        CampaignID
	actionID          ActionID
	triggeredAt       time.Time
	status            ActionTriggeredStatus
	statusDescription string
}

func (e *ActionTriggered) TriggeredAt() time.Time {
	return e.triggeredAt
}
