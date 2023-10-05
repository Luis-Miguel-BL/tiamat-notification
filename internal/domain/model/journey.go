package model

import (
	"fmt"
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

var AggregateTypeJourney = domain.AggregateType("journey")

type JourneyID string

type Journey struct {
	*domain.AggregateRoot
	journeyID       JourneyID
	workspaceID     WorkspaceID
	customerID      CustomerID
	campaignID      CampaignID
	currentActionID ActionID
	steps           map[ActionID]StepJourney
	startedAt       time.Time
	finishedAt      time.Time
}

type NewJourneyInput struct {
	WorkspaceID WorkspaceID
	CustomerID  CustomerID
	CampaignID  CampaignID
}

func NewJourney(input NewJourneyInput) (journey *Journey, err domain.DomainError) {
	journeyID := JourneyID(util.NewUUID())
	if input.WorkspaceID == "" {
		return journey, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	if input.CustomerID == "" {
		return journey, domain.NewInvalidEmptyParamError("CustomerID")
	}
	if input.CampaignID == "" {
		return journey, domain.NewInvalidEmptyParamError("CampaignID")
	}
	journey = &Journey{
		AggregateRoot: domain.NewAggregateRoot(AggregateTypeCustomer, domain.AggregateID(journeyID)),
		journeyID:     journeyID,
		workspaceID:   input.WorkspaceID,
		customerID:    input.CustomerID,
		campaignID:    input.CampaignID,
		steps:         make(map[ActionID]StepJourney),
		startedAt:     time.Now(),
	}

	return journey, nil
}

func (e *Journey) JourneyID() JourneyID {
	return e.journeyID
}
func (e *Journey) StartedAt() time.Time {
	return e.startedAt
}
func (e *Journey) FinishedAt() time.Time {
	return e.finishedAt
}
func (e *Journey) IsFinished() bool {
	return !e.finishedAt.IsZero()
}
func (e *Journey) CustomerID() CustomerID {
	return e.customerID
}
func (e *Journey) WorkspaceID() WorkspaceID {
	return e.workspaceID
}
func (e *Journey) CampaignID() CampaignID {
	return e.campaignID
}
func (e *Journey) CurrentActionID() ActionID {
	return e.currentActionID
}
func (e *Journey) StepJourney(actionID ActionID) (stepJourney StepJourney, found bool) {
	stepJourney, found = e.steps[actionID]
	return stepJourney, found
}

func (e *Journey) AppendNextStepJourney(actionID ActionID) (err error) {
	_, stepAlreadyExists := e.steps[actionID]
	if stepAlreadyExists {
		return domain.DomainError(fmt.Errorf("step journey already exist journey_id: %s action_id: %s", e.journeyID, actionID))
	}

	newStepJourney, err := newStepJourney(newStepJourneyInput{
		JourneyID: e.journeyID,
		ActionID:  actionID,
	})
	if err != nil {
		return err
	}
	e.steps[actionID] = *newStepJourney
	e.currentActionID = actionID

	e.AggregateRoot.AppendEvent(event.ActionTriggedEvent{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.ActionTriggedEventType,
			OccurredAt:    newStepJourney.TriggeredAt(),
			AggregateType: AggregateTypeJourney,
			AggregateID:   e.AggregateID(),
		}),
		CustomerID:    string(e.customerID),
		WorkspaceID:   string(e.workspaceID),
		CampaignID:    string(e.campaignID),
		ActionID:      string(actionID),
		StepJourneyID: string(newStepJourney.stepJourneyID),
		TriggeredAt:   newStepJourney.TriggeredAt(),
	})

	return nil
}

func (e *Journey) SkipStep(stepJourney StepJourney, reason event.SkippedReason) {
	stepJourney.status = StepJourneySkipped

	e.AggregateRoot.AppendEvent(event.StepJourneySkipped{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.StepJourneySkippedEventType,
			OccurredAt:    time.Now(),
			AggregateType: AggregateTypeJourney,
			AggregateID:   e.AggregateID(),
		}),
		CustomerID:    string(e.customerID),
		WorkspaceID:   string(e.workspaceID),
		CampaignID:    string(e.campaignID),
		ActionID:      string(stepJourney.ActionID()),
		StepJourneyID: string(stepJourney.StepJourneyID()),
		JourneyID:     string(e.journeyID),
		Reason:        reason,
		TriggeredAt:   stepJourney.TriggeredAt(),
	})
}
func (e *Journey) FinishSuccessfully(stepJourney StepJourney, nextActionID ActionID) {
	stepJourney.status = StepJourneyStatusSuccessed
	stepJourney.nextActionID = nextActionID

	e.AggregateRoot.AppendEvent(event.StepJourneySuccessed{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.StepJourneySuccessedEventType,
			OccurredAt:    time.Now(),
			AggregateType: AggregateTypeJourney,
			AggregateID:   e.AggregateID(),
		}),
		CustomerID:    string(e.customerID),
		WorkspaceID:   string(e.workspaceID),
		CampaignID:    string(e.campaignID),
		ActionID:      string(stepJourney.ActionID()),
		StepJourneyID: string(stepJourney.StepJourneyID()),
		JourneyID:     string(e.journeyID),
		TriggeredAt:   stepJourney.TriggeredAt(),
	})
}
func (e *Journey) FinishUnsuccessfully(stepJourney StepJourney, nextActionID ActionID, description string) {
	stepJourney.status = StepJourneyStatusSuccessed
	stepJourney.nextActionID = nextActionID

	e.AggregateRoot.AppendEvent(event.StepJourneyFailed{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.StepJourneyFailedEventType,
			OccurredAt:    time.Now(),
			AggregateType: AggregateTypeJourney,
			AggregateID:   e.AggregateID(),
		}),
		CustomerID:    string(e.customerID),
		WorkspaceID:   string(e.workspaceID),
		CampaignID:    string(e.campaignID),
		ActionID:      string(stepJourney.ActionID()),
		StepJourneyID: string(stepJourney.StepJourneyID()),
		JourneyID:     string(e.journeyID),
		Description:   description,
		TriggeredAt:   stepJourney.TriggeredAt(),
	})
}
