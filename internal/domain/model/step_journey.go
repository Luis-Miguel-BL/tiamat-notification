package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type StepJourneyID string

type StepJourneyStatus string

const (
	StepJourneyStatusTriggered StepJourneyStatus = "triggered"
	StepJourneyStatusScheduled StepJourneyStatus = "scheduled"
	StepJourneyStatusSuccessed StepJourneyStatus = "successed"
	StepJourneyStatusFailed    StepJourneyStatus = "failed"
	StepJourneySkipped         StepJourneyStatus = "skipped"
)

type StepJourney struct {
	stepJourneyID StepJourneyID
	workspaceID   WorkspaceID
	customerID    CustomerID
	campaignID    CampaignID
	actionID      ActionID
	nextActionID  ActionID
	triggeredAt   time.Time
	status        StepJourneyStatus
	trackingData  map[vo.Slug]vo.CustomAttributes
}

type NewStepJourneyInput struct {
	WorkspaceID WorkspaceID
	CustomerID  CustomerID
	CampaignID  CampaignID
	ActionID    ActionID
}

func NewStepJourney(input NewStepJourneyInput) (actionTriggered *StepJourney, err domain.DomainError) {
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
	actionTriggered = &StepJourney{
		stepJourneyID: StepJourneyID(util.NewUUID()),
		workspaceID:   input.WorkspaceID,
		customerID:    input.CustomerID,
		campaignID:    input.CampaignID,
		actionID:      input.ActionID,
		triggeredAt:   time.Now(),
		status:        StepJourneyStatusTriggered,
	}

	return actionTriggered, nil
}

func (e *StepJourney) StepJourneyID() StepJourneyID {
	return e.stepJourneyID
}
func (e *StepJourney) TriggeredAt() time.Time {
	return e.triggeredAt
}
func (e *StepJourney) CampaignID() CampaignID {
	return e.campaignID
}
func (e *StepJourney) ActionID() ActionID {
	return e.actionID
}
func (e *StepJourney) NextActionID() ActionID {
	return e.nextActionID
}

func (e *StepJourney) Finish(status StepJourneyStatus, nextActionID ActionID) {
	e.status = status
	e.nextActionID = nextActionID
}
func (e *StepJourney) AppendTrackingEvent(eventSlug vo.Slug, trackingData vo.CustomAttributes) {
	e.trackingData[eventSlug] = trackingData
}
