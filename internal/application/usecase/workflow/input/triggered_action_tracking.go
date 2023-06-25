package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type TriggeredActionTrackingInput struct {
	WorkspaceID   string
	StepJourneyID string
	CampaignID    string
	ActionID      string
	EventSlug     string
	TrackingData  map[string]any
}

func (c *TriggeredActionTrackingInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.StepJourneyID) {
		return errors.NewInvalidEmptyParamError("customer_journey_id")
	}
	if util.IsEmpty(c.CampaignID) {
		return errors.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.ActionID) {
		return errors.NewInvalidEmptyParamError("action_id")
	}
	if util.IsEmpty(c.EventSlug) {
		return errors.NewInvalidEmptyParamError("event_slug")
	}
	return nil
}
