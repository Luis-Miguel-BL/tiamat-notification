package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CustomerJourneyID string

type CustomerJourneyStatus string

const (
	CustomerJourneyStatusTriggered   CustomerJourneyStatus = "triggered"
	CustomerJourneyStatusScheduled   CustomerJourneyStatus = "scheduled"
	CustomerJourneyStatusSuccess     CustomerJourneyStatus = "success"
	CustomerJourneyStatusFailed      CustomerJourneyStatus = "failed"
	CustomerJourneyStatusFilterMatch CustomerJourneyStatus = "filter-matched"
)

type CustomerJourney struct {
	customerJourneyID CustomerJourneyID
	workspaceID       WorkspaceID
	customerID        CustomerID
	campaignID        CampaignID
	actionID          ActionID
	nextActionID      ActionID
	triggeredAt       time.Time
	status            CustomerJourneyStatus
	trackingData      map[vo.Slug]vo.CustomAttributes
}

type NewCustomerJourneyInput struct {
	WorkspaceID WorkspaceID
	CustomerID  CustomerID
	CampaignID  CampaignID
	ActionID    ActionID
}

func NewCustomerJourney(input NewCustomerJourneyInput) (actionTriggered *CustomerJourney, err domain.DomainError) {
	if input.WorkspaceID == "" {
		return actionTriggered, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	if input.CustomerID == "" {
		return actionTriggered, domain.NewInvalidEmptyParamError("CustomerID")
	}
	if input.CampaignID == "" {
		return actionTriggered, domain.NewInvalidEmptyParamError("CampaignID")
	}
	if input.ActionID == "" {
		return actionTriggered, domain.NewInvalidEmptyParamError("ActionID")
	}
	actionTriggered = &CustomerJourney{
		customerJourneyID: CustomerJourneyID(util.NewUUID()),
		workspaceID:       input.WorkspaceID,
		customerID:        input.CustomerID,
		campaignID:        input.CampaignID,
		actionID:          input.ActionID,
		triggeredAt:       time.Now(),
		status:            CustomerJourneyStatusTriggered,
	}

	return actionTriggered, nil
}

func (e *CustomerJourney) CustomerJourneyID() CustomerJourneyID {
	return e.customerJourneyID
}
func (e *CustomerJourney) TriggeredAt() time.Time {
	return e.triggeredAt
}
func (e *CustomerJourney) CampaignID() CampaignID {
	return e.campaignID
}
func (e *CustomerJourney) ActionID() ActionID {
	return e.actionID
}
func (e *CustomerJourney) NextActionID() ActionID {
	return e.nextActionID
}

func (e *CustomerJourney) Finish(status CustomerJourneyStatus, nextActionID ActionID) {
	e.status = status
	e.nextActionID = nextActionID
}
func (e *CustomerJourney) AppendTrackingEvent(eventSlug vo.Slug, trackingData vo.CustomAttributes) {
	e.trackingData[eventSlug] = trackingData
}
