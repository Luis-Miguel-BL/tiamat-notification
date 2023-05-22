package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

var CampaignFilterMatchedType = domain.EventType("campaign-filter-mached")

type CampaignFilterMatched struct {
	*domain.DomainEventBase
	CustomerID        string
	WorkspaceID       string
	CampaignID        string
	ActionID          string
	CustomerJourneyID string
	TriggeredAt       time.Time
}
