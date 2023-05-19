package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateWorkspaceCommand struct {
	Slug string `json:"slug,omitempty"`
}

func (c *CreateWorkspaceCommand) Validate() (err error) {
	if util.IsEmpty(c.Slug) {
		return application.NewInvalidEmptyParamError("slug")
	}
	return nil
}

type UpdateWorkspaceCommand struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
	Slug        string `json:"slug,omitempty"`
}

func (c *UpdateWorkspaceCommand) Validate() (err error) {
	if util.IsEmpty(c.Slug) {
		return application.NewInvalidEmptyParamError("slug")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return application.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type DeleteWorkspaceCommand struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *DeleteWorkspaceCommand) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return application.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type GetWorkspaceCommand struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *GetWorkspaceCommand) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return application.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
