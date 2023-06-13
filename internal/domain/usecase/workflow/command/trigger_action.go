package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type TriggerActionCommand struct {
	WorkspaceID   string `json:"workspace_id,omitempty"`
	CustomerID    string `json:"customer_id,omitempty"`
	CampaignID    string `json:"campaign_id,omitempty"`
	ActionID      string `json:"action_id,omitempty"`
	StepJourneyID string `json:"step_journey_id,omitempty"`
	JourneyID     string `json:"journey_id,omitempty"`
}

func (c *TriggerActionCommand) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.CustomerID) {
		return domain.NewInvalidEmptyParamError("customer_id")
	}
	if util.IsEmpty(c.CampaignID) {
		return domain.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.ActionID) {
		return domain.NewInvalidEmptyParamError("action_id")
	}
	if util.IsEmpty(c.StepJourneyID) {
		return domain.NewInvalidEmptyParamError("step_journey_id")
	}
	if util.IsEmpty(c.JourneyID) {
		return domain.NewInvalidEmptyParamError("journey_id")
	}
	return nil
}
