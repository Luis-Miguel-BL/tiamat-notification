package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateCustomerEventCommand struct {
	CustomerID       string         `json:"customer_id,omitempty"`
	WorkspaceID      string         `json:"workspace_id,omitempty"`
	CustomerEventID  string         `json:"customer_event_id,omitempty"`
	Slug             string         `json:"slug,omitempty"`
	CustomAttributes map[string]any `json:"custom_attributes,omitempty"`
}

func (c *CreateCustomerEventCommand) Validate() (err error) {
	if util.IsEmpty(c.CustomerEventID) {
		c.CustomerEventID = util.NewUUID()
	}
	if util.IsEmpty(c.CustomerID) {
		return domain.NewInvalidEmptyParamError("customer_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Slug) {
		return domain.NewInvalidEmptyParamError("slug")
	}
	return nil
}
