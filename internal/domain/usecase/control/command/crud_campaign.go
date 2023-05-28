package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateCampaignCommand struct {
	WorkspaceID             string   `json:"workspace_id,omitempty"`
	Slug                    string   `json:"slug,omitempty"`
	RetriggerDelayInSeconds int      `json:"retrigger_delay_in_seconds,omitempty"`
	Triggers                []string `json:"triggers,omitempty"`
	Filters                 []string `json:"filters,omitempty"`
}

func (c *CreateCampaignCommand) Validate() (err error) {
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

type UpdateCampaignCommand struct {
	WorkspaceID             string   `json:"workspace_id,omitempty"`
	CampaignID              string   `json:"campaign_id,omitempty"`
	Slug                    string   `json:"slug,omitempty"`
	RetriggerDelayInSeconds int      `json:"retrigger_delay_in_seconds,omitempty"`
	Triggers                []string `json:"triggers,omitempty"`
	Filters                 []string `json:"filters,omitempty"`
}

func (c *UpdateCampaignCommand) Validate() (err error) {
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

type DeleteCampaignCommand struct {
	CampaignID  string `json:"campaign_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *DeleteCampaignCommand) Validate() (err error) {
	if util.IsEmpty(c.CampaignID) {
		return domain.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type GetCampaignCommand struct {
	CampaignID  string `json:"campaign_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *GetCampaignCommand) Validate() (err error) {
	if util.IsEmpty(c.CampaignID) {
		return domain.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type ListCampaignCommand struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *ListCampaignCommand) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
