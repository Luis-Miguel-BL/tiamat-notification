package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateWorkspaceInput struct {
	Slug string `json:"slug,omitempty"`
}

func (c *CreateWorkspaceInput) Validate() (err error) {
	if util.IsEmpty(c.Slug) {
		return domain.NewInvalidEmptyParamError("slug")
	}
	return nil
}

type UpdateWorkspaceInput struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
	Slug        string `json:"slug,omitempty"`
}

func (c *UpdateWorkspaceInput) Validate() (err error) {
	if util.IsEmpty(c.Slug) {
		return domain.NewInvalidEmptyParamError("slug")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type DeleteWorkspaceInput struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *DeleteWorkspaceInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type GetWorkspaceInput struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *GetWorkspaceInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
