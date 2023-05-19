package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type ActionID string

type ActionType string

const (
	ActionTypeNotification ActionType = "notification"
	ActionTypeDelay        ActionType = "delay"
	ActionTypeFlow         ActionType = "flow"
)

type Action struct {
	actionID     ActionID
	campaignID   CampaignID
	slug         vo.Slug
	actionType   ActionType
	nextActionID ActionID
	behaviorType BehaviorType
	behavior     ActionBehavior
	createdAt    time.Time
	updatedAt    time.Time
}

func (e *Action) CampaignID() CampaignID {
	return e.campaignID
}
func (e *Action) ActionID() ActionID {
	return e.actionID
}
func (e *Action) BehaviorType() BehaviorType {
	return e.behaviorType
}
