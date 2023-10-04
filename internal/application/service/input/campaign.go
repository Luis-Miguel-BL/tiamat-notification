package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateCampaignInput struct {
	WorkspaceID             string
	Slug                    string
	RetriggerDelayInSeconds int
	Triggers                []string
	Filters                 []string
}

func (c *CreateCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Slug) {
		return errors.NewInvalidEmptyParamError("slug")
	}
	if len(c.Triggers) == 0 {
		return errors.NewInvalidEmptyParamError("triggers")
	}
	return nil
}

type UpdateCampaignInput struct {
	WorkspaceID             string
	CampaignID              string
	Slug                    string
	RetriggerDelayInSeconds int
	Triggers                []string
	Filters                 []string
}

func (c *UpdateCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.CampaignID) {
		return errors.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Slug) {
		return errors.NewInvalidEmptyParamError("slug")
	}
	if len(c.Triggers) == 0 {
		return errors.NewInvalidEmptyParamError("triggers")
	}
	return nil
}

type DeleteCampaignInput struct {
	CampaignID  string
	WorkspaceID string
}

func (c *DeleteCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.CampaignID) {
		return errors.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type GetCampaignInput struct {
	CampaignID  string
	WorkspaceID string
}

func (c *GetCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.CampaignID) {
		return errors.NewInvalidEmptyParamError("campaign_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type ListCampaignInput struct {
	WorkspaceID string
}

func (c *ListCampaignInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
