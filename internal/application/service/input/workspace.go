package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateWorkspaceInput struct {
	Slug string
}

func (c *CreateWorkspaceInput) Validate() (err error) {
	if util.IsEmpty(c.Slug) {
		return errors.NewInvalidEmptyParamError("slug")
	}
	return nil
}

type UpdateWorkspaceInput struct {
	WorkspaceID string
	Slug        string
}

func (c *UpdateWorkspaceInput) Validate() (err error) {
	if util.IsEmpty(c.Slug) {
		return errors.NewInvalidEmptyParamError("slug")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type DeleteWorkspaceInput struct {
	WorkspaceID string
}

func (c *DeleteWorkspaceInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type GetWorkspaceInput struct {
	WorkspaceID string
}

func (c *GetWorkspaceInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
