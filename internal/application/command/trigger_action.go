package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type TriggerActionCommand struct {
	WorkspaceID       string `json:"workspace_id,omitempty"`
	CustomerID        string `json:"customer_id,omitempty"`
	CampaignID        string `json:"campaign_id,omitempty"`
	ActionID          string `json:"action_id,omitempty"`
	ActionTriggeredID string `json:"action_triggered_id,omitempty"`
}

func (c *TriggerActionCommand) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return application.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.CustomerID) {
		return application.NewInvalidEmptyParamError("customer_id")
	}
	if util.IsEmpty(c.CampaignID) {
		return application.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.ActionID) {
		return application.NewInvalidEmptyParamError("action_id")
	}
	if util.IsEmpty(c.ActionTriggeredID) {
		return application.NewInvalidEmptyParamError("action_triggered_id")
	}
	return nil
}