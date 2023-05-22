package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type ActionID string

type ActionType string

const (
	ActionTypeNotification ActionType = "notification"
	ActionTypeDelay        ActionType = "delay"
	ActionTypeFlow         ActionType = "flow"
)

var AvailableActionType map[ActionType]struct{} = map[ActionType]struct{}{
	ActionTypeNotification: {},
	ActionTypeDelay:        {},
	ActionTypeFlow:         {},
}

type Action struct {
	actionID     ActionID
	slug         vo.Slug
	actionType   ActionType
	nextActionID ActionID
	behaviorType BehaviorType
	behavior     ActionBehavior
	createdAt    time.Time
	updatedAt    time.Time
}

type NewActionInput struct {
	ActionID     ActionID
	Slug         vo.Slug
	ActionType   ActionType
	NextActionID ActionID
	BehaviorType BehaviorType
	Behavior     ActionBehavior
}

func NewAction(input NewActionInput) (action Action, err domain.DomainError) {
	if util.IsEmpty(string(input.ActionID)) {
		return action, domain.NewInvalidEmptyParamError("action_id")
	}
	if util.IsEmpty(string(input.Slug)) {
		return action, domain.NewInvalidEmptyParamError("slug")
	}

	_, availableActionType := AvailableActionType[input.ActionType]
	if !availableActionType {
		return action, domain.NewInvalidParamError("action_type")
	}
	_, availableBehaviorType := AvailableBehaviorType[input.BehaviorType]
	if !availableBehaviorType {
		return action, domain.NewInvalidParamError("behavior_type")
	}

	return Action{
		actionID:     input.ActionID,
		slug:         input.Slug,
		actionType:   input.ActionType,
		nextActionID: input.NextActionID,
		behaviorType: input.BehaviorType,
		behavior:     input.Behavior,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}, nil
}

func (e *Action) ActionID() ActionID {
	return e.actionID
}
func (e *Action) BehaviorType() BehaviorType {
	return e.behaviorType
}
