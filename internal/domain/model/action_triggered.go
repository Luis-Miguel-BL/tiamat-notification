package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type ActionTriggeredID string

func NewActionTriggeredID(actionTriggeredID string) ActionTriggeredID {
	return ActionTriggeredID(actionTriggeredID)
}

type ActionTriggeredStatus string

const (
	ActionTriggeredStatusTriggered   ActionTriggeredStatus = "triggered"
	ActionTriggeredStatusScheduled   ActionTriggeredStatus = "scheduled"
	ActionTriggeredStatusSuccess     ActionTriggeredStatus = "success"
	ActionTriggeredStatusFailed      ActionTriggeredStatus = "failed"
	ActionTriggeredStatusFilterMatch ActionTriggeredStatus = "filter-matched"
)

type ActionTriggered struct {
	actionTriggeredID ActionTriggeredID
	workspaceID       WorkspaceID
	customerID        CustomerID
	campaignID        CampaignID
	actionID          ActionID
	triggeredAt       time.Time
	status            ActionTriggeredStatus
}

type NewActionTriggeredInput struct {
	WorkspaceID WorkspaceID
	CustomerID  CustomerID
	CampaignID  CampaignID
	ActionID    ActionID
}

func NewActionTriggered(input NewActionTriggeredInput) (actionTriggered *ActionTriggered, err domain.DomainError) {
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
	actionTriggered = &ActionTriggered{
		actionTriggeredID: ActionTriggeredID(util.NewUUID()),
		workspaceID:       input.WorkspaceID,
		customerID:        input.CustomerID,
		campaignID:        input.CampaignID,
		actionID:          input.ActionID,
		triggeredAt:       time.Now(),
		status:            ActionTriggeredStatusTriggered,
	}

	return actionTriggered, nil
}

func (e *ActionTriggered) ActionTriggeredID() ActionTriggeredID {
	return e.actionTriggeredID
}
func (e *ActionTriggered) TriggeredAt() time.Time {
	return e.triggeredAt
}
func (e *ActionTriggered) CampaignID() CampaignID {
	return e.campaignID
}
func (e *ActionTriggered) ActionID() ActionID {
	return e.actionID
}

func (e *ActionTriggered) SetStatus(status ActionTriggeredStatus) {
	e.status = status
}
