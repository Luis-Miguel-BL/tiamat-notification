package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateSegmentInput struct {
	Slug        string
	WorkspaceID string
	Conditions  []Condition
}
type Condition struct {
	ConditionTarget string
	ConditionType   string
	EventSlug       string
	AttributeKey    string
	AttributeValue  any
}

func (c *CreateSegmentInput) Validate() (err error) {
	if util.IsEmpty(c.Slug) {
		return errors.NewInvalidEmptyParamError("slug")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if len(c.Conditions) == 0 {
		return errors.NewInvalidEmptyParamError("conditions")
	}
	for _, condition := range c.Conditions {
		if util.IsEmpty(condition.ConditionTarget) {
			return errors.NewInvalidEmptyParamError("condition_target")
		}
		if util.IsEmpty(condition.ConditionType) {
			return errors.NewInvalidEmptyParamError("condition_type")
		}
	}
	return nil
}

type UpdateSegmentInput struct {
	SegmentID   string
	Slug        string
	WorkspaceID string
	Conditions  []Condition
}

func (c *UpdateSegmentInput) Validate() (err error) {
	if util.IsEmpty(c.SegmentID) {
		return errors.NewInvalidEmptyParamError("segment_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Slug) {
		return errors.NewInvalidEmptyParamError("slug")
	}
	if len(c.Conditions) == 0 {
		return errors.NewInvalidEmptyParamError("conditions")
	}
	for _, condition := range c.Conditions {
		if util.IsEmpty(condition.ConditionTarget) {
			return errors.NewInvalidEmptyParamError("condition_target")
		}
		if util.IsEmpty(condition.ConditionType) {
			return errors.NewInvalidEmptyParamError("condition_type")
		}
	}
	return nil
}

type DeleteSegmentInput struct {
	SegmentID   string
	WorkspaceID string
}

func (c *DeleteSegmentInput) Validate() (err error) {
	if util.IsEmpty(c.SegmentID) {
		return errors.NewInvalidEmptyParamError("segment_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type GetSegmentInput struct {
	SegmentID   string
	WorkspaceID string
}

func (c *GetSegmentInput) Validate() (err error) {
	if util.IsEmpty(c.SegmentID) {
		return errors.NewInvalidEmptyParamError("segment_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type ListSegmentInput struct {
	WorkspaceID string
}

func (c *ListSegmentInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
