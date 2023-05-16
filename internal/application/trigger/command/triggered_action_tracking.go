package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type TriggeredActionTrackingCommand struct {
	WorkspaceID       string         `json:"workspace_id,omitempty"`
	ActionTriggeredID string         `json:"action_triggered_id,omitempty"`
	CampaignID        string         `json:"campaign_id,omitempty"`
	ActionID          string         `json:"action_id,omitempty"`
	EventSlug         string         `json:"event_slug,omitempty"`
	TrackingData      map[string]any `json:"tracking_data,omitempty"`
}

func (c *TriggeredActionTrackingCommand) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return application.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.ActionTriggeredID) {
		return application.NewInvalidEmptyParamError("action_triggered_id")
	}
	if util.IsEmpty(c.CampaignID) {
		return application.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.ActionID) {
		return application.NewInvalidEmptyParamError("action_id")
	}
	if util.IsEmpty(c.EventSlug) {
		return application.NewInvalidEmptyParamError("event_slug")
	}
	return nil
}
