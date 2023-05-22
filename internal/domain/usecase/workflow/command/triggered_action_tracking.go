package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type TriggeredActionTrackingCommand struct {
	WorkspaceID       string         `json:"workspace_id,omitempty"`
	CustomerJourneyID string         `json:"customer_journey_id,omitempty"`
	CampaignID        string         `json:"campaign_id,omitempty"`
	ActionID          string         `json:"action_id,omitempty"`
	EventSlug         string         `json:"event_slug,omitempty"`
	TrackingData      map[string]any `json:"tracking_data,omitempty"`
}

func (c *TriggeredActionTrackingCommand) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.CustomerJourneyID) {
		return domain.NewInvalidEmptyParamError("customer_journey_id")
	}
	if util.IsEmpty(c.CampaignID) {
		return domain.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.ActionID) {
		return domain.NewInvalidEmptyParamError("action_id")
	}
	if util.IsEmpty(c.EventSlug) {
		return domain.NewInvalidEmptyParamError("event_slug")
	}
	return nil
}
