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
	journeyID     JourneyID
	stepJourneyID StepJourneyID
	actionID      ActionID
	nextActionID  ActionID
	triggeredAt   time.Time
	status        StepJourneyStatus
	trackingData  map[vo.Slug]vo.CustomAttributes
}

type newStepJourneyInput struct {
	JourneyID JourneyID
	ActionID  ActionID
}

func newStepJourney(input newStepJourneyInput) (stepJourney *StepJourney, err domain.DomainError) {
	if input.JourneyID == "" {
		return stepJourney, domain.NewInvalidEmptyParamError("JourneyID")
	}
	stepJourney = &StepJourney{
		stepJourneyID: StepJourneyID(util.NewUUID()),
		journeyID:     input.JourneyID,
		actionID:      input.ActionID,
		triggeredAt:   time.Now(),
		status:        StepJourneyStatusTriggered,
	}

	return stepJourney, nil
}

func (e *StepJourney) StepJourneyID() StepJourneyID {
	return e.stepJourneyID
}
func (e *StepJourney) TriggeredAt() time.Time {
	return e.triggeredAt
}
func (e *StepJourney) JourneyID() JourneyID {
	return e.journeyID
}
func (e *StepJourney) ActionID() ActionID {
	return e.actionID
}
func (e *StepJourney) NextActionID() ActionID {
	return e.nextActionID
}

func (e *StepJourney) AppendTrackingEvent(eventSlug vo.Slug, trackingData vo.CustomAttributes) {
	e.trackingData[eventSlug] = trackingData
}
