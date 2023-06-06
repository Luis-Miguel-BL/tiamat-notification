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
	campaignID            CampaignID
	workspaceID           WorkspaceID
	slug                  vo.Slug
	isActive              bool
	retriggerDelay        time.Duration
	actions               map[ActionID]Action
	firstActionID         ActionID
	triggers              []SegmentID
	filters               []SegmentID
	notificationTimeRange vo.TimeRange
	createdAt             time.Time
	updatedAt             time.Time
}

type NewCampaignInput struct {
	WorkspaceID    WorkspaceID
	Slug           vo.Slug
	RetriggerDelay time.Duration
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
	if len(input.Triggers) == 0 {
		return segment, domain.NewInvalidEmptyParamError("Triggers")
	}
	return &Campaign{
		AggregateRoot:  domain.NewAggregateRoot(AggregateTypeCampaign, domain.AggregateID(campaignID)),
		campaignID:     campaignID,
		workspaceID:    input.WorkspaceID,
		slug:           input.Slug,
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

func (e *Campaign) SetSlug(slug vo.Slug) {
	e.slug = slug
	e.updatedAt = time.Now()
}
func (e *Campaign) Triggers() []SegmentID {
	return e.triggers
}
func (e *Campaign) SetTriggers(triggers []SegmentID) {
	e.triggers = triggers
	e.updatedAt = time.Now()
}
func (e *Campaign) Filters() []SegmentID {
	return e.filters
}
func (e *Campaign) SetFilters(filters []SegmentID) {
	e.filters = filters
	e.updatedAt = time.Now()
}
func (e *Campaign) RetriggerDelay() time.Duration {
	return e.retriggerDelay
}
func (e *Campaign) SetRetriggerDelay(retriggerDelay time.Duration) {
	e.retriggerDelay = retriggerDelay
	e.updatedAt = time.Now()
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
func (e *Campaign) SetActions(firstActionID ActionID, actions map[ActionID]Action) {
	e.firstActionID = firstActionID
	e.actions = actions
}
func (e *Campaign) IsActive() bool {
	return e.isActive
}
func (e *Campaign) NextAvailableTimeToTriggerNotification(currentTime time.Time) (alreadyAvailable bool, nextTime time.Time) {
	if e.notificationTimeRange.IsAvailable(currentTime) {
		return true, nextTime
	}
	return false, e.notificationTimeRange.NextAvailableTime(currentTime)
}
