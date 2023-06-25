package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type Action struct {
	ActionID      string
	Slug          string
	ActionType    string
	NextActionsID []string
	BehaviorType  string
	Behavior      model.ActionBehavior
}
type SaveActionsInput struct {
	WorkspaceID   string
	CampaignID    string
	Actions       []Action
	FirstActionID string
}

func (c *SaveActionsInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if len(c.Actions) == 0 {
		return errors.NewInvalidEmptyParamError("actions")
	}
	if c.FirstActionID == "" {
		return errors.NewInvalidEmptyParamError("first_action_id")
	}
	for _, action := range c.Actions {
		if util.IsEmpty(action.Slug) {
			return errors.NewInvalidEmptyParamError("action.slug")
		}
		if util.IsEmpty(action.ActionType) {
			return errors.NewInvalidEmptyParamError("action.type")
		}
		if util.IsEmpty(action.BehaviorType) {
			return errors.NewInvalidEmptyParamError("action.behabior_type")
		}
	}
	return nil
}
