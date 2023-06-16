package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateCampaignInput struct {
	WorkspaceID             string   `json:"workspace_id,omitempty"`
	Slug                    string   `json:"slug,omitempty"`
	RetriggerDelayInSeconds int      `json:"retrigger_delay_in_seconds,omitempty"`
	Triggers                []string `json:"triggers,omitempty"`
	Filters                 []string `json:"filters,omitempty"`
}

func (c *CreateCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Slug) {
		return domain.NewInvalidEmptyParamError("slug")
	}
	if len(c.Triggers) == 0 {
		return domain.NewInvalidEmptyParamError("triggers")
	}
	return nil
}

type UpdateCampaignInput struct {
	WorkspaceID             string   `json:"workspace_id,omitempty"`
	CampaignID              string   `json:"campaign_id,omitempty"`
	Slug                    string   `json:"slug,omitempty"`
	RetriggerDelayInSeconds int      `json:"retrigger_delay_in_seconds,omitempty"`
	Triggers                []string `json:"triggers,omitempty"`
	Filters                 []string `json:"filters,omitempty"`
}

func (c *UpdateCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.CampaignID) {
		return domain.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Slug) {
		return domain.NewInvalidEmptyParamError("slug")
	}
	if len(c.Triggers) == 0 {
		return domain.NewInvalidEmptyParamError("triggers")
	}
	return nil
}

type DeleteCampaignInput struct {
	CampaignID  string `json:"campaign_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *DeleteCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.CampaignID) {
		return domain.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type GetCampaignInput struct {
	CampaignID  string `json:"campaign_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *GetCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.CampaignID) {
		return domain.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type ListCampaignInput struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *ListCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
