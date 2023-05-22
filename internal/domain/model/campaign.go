package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

var AggregateTypeCampaign = domain.AggregateType("campaign")

type CampaignID string

type Campaign struct {
	*domain.AggregateRoot
	campaignID     CampaignID
	workspaceID    WorkspaceID
	slug           vo.Slug
	isActive       bool
	retriggerDelay time.Duration
	actions        map[ActionID]Action
	firstActionID  ActionID
	triggers       []SegmentID
	filters        []SegmentID
	createdAt      time.Time
	updatedAt      time.Time
}

type NewCampaignInput struct {
	WorkspaceID    WorkspaceID
	Slug           vo.Slug
	RetriggerDelay time.Duration
	Actions        map[ActionID]Action
	FirstActionID  ActionID
	Triggers       []SegmentID
	Filters        []SegmentID
}

func NewCampaign(input NewCampaignInput) (segment *Campaign, err domain.DomainError) {
	campaignID := CampaignID(util.NewUUID())
	if input.WorkspaceID == "" {
		return segment, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	if input.Slug == "" {
		return segment, domain.NewInvalidEmptyParamError("Slug")
	}
	if len(input.Actions) == 0 {
		return segment, domain.NewInvalidEmptyParamError("Actions")
	}
	if input.FirstActionID == "" {
		return segment, domain.NewInvalidEmptyParamError("FirstActionID")
	}
	if len(input.Triggers) == 0 {
		return segment, domain.NewInvalidEmptyParamError("Triggers")
	}
	return &Campaign{
		AggregateRoot:  domain.NewAggregateRoot(AggregateTypeCampaign, domain.AggregateID(campaignID)),
		campaignID:     campaignID,
		workspaceID:    input.WorkspaceID,
		slug:           input.Slug,
		actions:        input.Actions,
		firstActionID:  input.FirstActionID,
		retriggerDelay: input.RetriggerDelay,
		triggers:       input.Triggers,
		filters:        input.Filters,
		isActive:       true,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
	}, nil
}

func (e *Campaign) CampaignID() CampaignID {
	return e.campaignID
}

func (e *Campaign) Triggers() []SegmentID {
	return e.triggers
}
func (e *Campaign) Filters() []SegmentID {
	return e.filters
}
func (e *Campaign) RetriggerDelay() time.Duration {
	return e.retriggerDelay
}
func (e *Campaign) FirstActionID() ActionID {
	return e.firstActionID
}
func (e *Campaign) Actions() map[ActionID]Action {
	return e.actions
}
func (e *Campaign) Action(actionID ActionID) (action Action, err error) {
	action, found := e.actions[actionID]
	if !found {
		return action, domain.NewNotFoundError("action")
	}
	return action, nil
}

func (e *Campaign) MustBeTriggered(lastTriggeredDate time.Time) bool {
	if lastTriggeredDate.IsZero() {
		return true
	}
	if lastTriggeredDate.Add(e.retriggerDelay).After(time.Now()) {
		return true
	}
	return false
}
