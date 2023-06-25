package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type TriggerActionInput struct {
	WorkspaceID   string
	CustomerID    string
	CampaignID    string
	ActionID      string
	StepJourneyID string
	JourneyID     string
}

func (c *TriggerActionInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.CustomerID) {
		return errors.NewInvalidEmptyParamError("customer_id")
	}
	if util.IsEmpty(c.CampaignID) {
		return errors.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.ActionID) {
		return errors.NewInvalidEmptyParamError("action_id")
	}
	if util.IsEmpty(c.StepJourneyID) {
		return errors.NewInvalidEmptyParamError("step_journey_id")
	}
	if util.IsEmpty(c.JourneyID) {
		return errors.NewInvalidEmptyParamError("journey_id")
	}
	return nil
}
