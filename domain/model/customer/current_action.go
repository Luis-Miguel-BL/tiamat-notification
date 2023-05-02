package customer

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain/model/campaign"
)

type CurrentActionID string
type CurrentActionStatus string

const (
	CurrentActionStatusSuccess   CurrentActionStatus = "success"
	CurrentActionStatusScheduled CurrentActionStatus = "scheduled"
	CurrentActionStatusFailed    CurrentActionStatus = "failed"
)

type CurrentAction struct {
	CurrentActionID   CurrentActionID
	CampaignID        campaign.CampaignID
	ActionID          campaign.ActionID
	TriggeredAt       time.Time
	ScheduledFor      time.Time
	Status            CurrentActionStatus
	StatusDescription string
}
