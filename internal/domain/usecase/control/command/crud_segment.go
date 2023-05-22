package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateSegmentCommand struct {
	Slug        string      `json:"slug,omitempty"`
	WorkspaceID string      `json:"workspace_id,omitempty"`
	Conditions  []Condition `json:"conditions,omitempty"`
}
type Condition struct {
	ConditionTarget string `json:"condition_target,omitempty"`
	ConditionType   string `json:"condition_type,omitempty"`
	EventSlug       string `json:"event_slug,omitempty"`
	AttributeKey    string `json:"attribute_key,omitempty"`
	AttributeValue  any    `json:"attribute_value,omitempty"`
}

func (c *CreateSegmentCommand) Validate() (err error) {
	if util.IsEmpty(c.Slug) {
		return domain.NewInvalidEmptyParamError("slug")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if len(c.Conditions) == 0 {
		return domain.NewInvalidEmptyParamError("conditions")
	}
	for _, condition := range c.Conditions {
		if util.IsEmpty(condition.ConditionTarget) {
			return domain.NewInvalidEmptyParamError("condition_target")
		}
		if util.IsEmpty(condition.ConditionType) {
			return domain.NewInvalidEmptyParamError("condition_type")
		}
	}
	return nil
}

type UpdateSegmentCommand struct {
	SegmentID   string      `json:"segment_id,omitempty"`
	Slug        string      `json:"slug,omitempty"`
	WorkspaceID string      `json:"workspace_id,omitempty"`
	Conditions  []Condition `json:"conditions,omitempty"`
}

func (c *UpdateSegmentCommand) Validate() (err error) {
	if util.IsEmpty(c.SegmentID) {
		return domain.NewInvalidEmptyParamError("segment_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Slug) {
		return domain.NewInvalidEmptyParamError("slug")
	}
	if len(c.Conditions) == 0 {
		return domain.NewInvalidEmptyParamError("conditions")
	}
	for _, condition := range c.Conditions {
		if util.IsEmpty(condition.ConditionTarget) {
			return domain.NewInvalidEmptyParamError("condition_target")
		}
		if util.IsEmpty(condition.ConditionType) {
			return domain.NewInvalidEmptyParamError("condition_type")
		}
	}
	return nil
}

type DeleteSegmentCommand struct {
	SegmentID   string `json:"segment_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *DeleteSegmentCommand) Validate() (err error) {
	if util.IsEmpty(c.SegmentID) {
		return domain.NewInvalidEmptyParamError("segment_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type GetSegmentCommand struct {
	SegmentID   string `json:"segment_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *GetSegmentCommand) Validate() (err error) {
	if util.IsEmpty(c.SegmentID) {
		return domain.NewInvalidEmptyParamError("segment_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}

type ListSegmentCommand struct {
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *ListSegmentCommand) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
