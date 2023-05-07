package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model/action_behavior"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type ActionID string
type Action struct {
	ActionID     ActionID
	Slug         vo.Slug
	Type         ActionType
	NextActionID ActionID
	Behavior     ActionBehavior
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ActionType string

const (
	ActionTypeNotification ActionType = "notification"
	ActionTypeDelay        ActionType = "delay"
	ActionTypeFlow         ActionType = "flow"
)

type ActionBehavior struct {
	SendEmail    action_behavior.SendEmail
	SendSMS      action_behavior.SendSMS
	SendWhatsApp action_behavior.SendWhatsApp
	WaitFor      action_behavior.WaitFor
	WaitUntil    action_behavior.WaitUntil
	IfAttribute  action_behavior.IfAttribute
	Random       action_behavior.Random
	Split        action_behavior.Split
}
