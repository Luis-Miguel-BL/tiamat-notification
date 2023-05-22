package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateCampaignCommand struct {
	WorkspaceID             string   `json:"workspace_id,omitempty"`
	Slug                    string   `json:"slug,omitempty"`
	RetriggerDelayInSeconds int      `json:"retrigger_delay_in_seconds,omitempty"`
	Actions                 []Action `json:"actions,omitempty"`
	FirstActionID           string   `json:"c,omitempty"`
	Triggers                []string `json:"triggers,omitempty"`
	Filters                 []string `json:"filters,omitempty"`
}

type Action struct {
	ActionID     string
	Slug         string
	ActionType   string
	NextActionID string
	BehaviorType string
	Behavior     model.ActionBehavior
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

type UpdateCampaignCommand struct {
	CampaignID  string      `json:"campaign_id,omitempty"`
	Slug        string      `json:"slug,omitempty"`
	WorkspaceID string      `json:"workspace_id,omitempty"`
	Conditions  []Condition `json:"conditions,omitempty"`
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
