package customer

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain/model/campaign"
)

type ActionsTriggeredID string
type ActionsTriggeredStatus string

const (
	ActionsTriggeredStatusSuccess   ActionsTriggeredStatus = "success"
	ActionsTriggeredStatusScheduled ActionsTriggeredStatus = "scheduled"
	ActionsTriggeredStatusFailed    ActionsTriggeredStatus = "failed"
)

type ActionsTriggered struct {
	ActionsTriggeredID ActionsTriggeredID
	CampaignID         campaign.CampaignID
	ActionID           campaign.ActionID
	TriggeredAt        time.Time
	Status             ActionsTriggeredStatus
	StatusDescription  string
}
