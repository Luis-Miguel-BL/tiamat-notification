package model

import (
	"time"
)

type ActionTriggeredID string
type ActionTriggeredStatus string

const (
	ActionTriggeredStatusSuccess   ActionTriggeredStatus = "success"
	ActionTriggeredStatusScheduled ActionTriggeredStatus = "scheduled"
	ActionTriggeredStatusFailed    ActionTriggeredStatus = "failed"
)

type ActionTriggered struct {
	ActionTriggeredID ActionTriggeredID
	CampaignID        CampaignID
	ActionID          ActionID
	TriggeredAt       time.Time
	Status            ActionTriggeredStatus
	StatusDescription string
}
