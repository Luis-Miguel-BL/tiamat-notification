package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
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
type SaveActionsCommand struct {
	WorkspaceID   string   `json:"workspace_id,omitempty"`
	CampaignID    string   `json:"campaign_id,omitempty"`
	Actions       []Action `json:"actions,omitempty"`
	FirstActionID string   `json:"first_action_ids,omitempty"`
}

func (c *SaveActionsCommand) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if len(c.Actions) == 0 {
		return domain.NewInvalidEmptyParamError("actions")
	}
	if c.FirstActionID == "" {
		return domain.NewInvalidEmptyParamError("first_action_id")
	}
	for _, action := range c.Actions {
		if util.IsEmpty(action.Slug) {
			return domain.NewInvalidEmptyParamError("action.slug")
		}
		if util.IsEmpty(action.ActionType) {
			return domain.NewInvalidEmptyParamError("action.type")
		}
		if util.IsEmpty(action.BehaviorType) {
			return domain.NewInvalidEmptyParamError("action.behabior_type")
		}
	}
	return nil
}
