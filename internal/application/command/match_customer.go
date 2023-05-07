package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type MatchCustomerCommand struct {
	CustomerID  string `json:"customer_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *MatchCustomerCommand) Validate() (err error) {
	if util.IsEmpty(c.CustomerID) {
		return application.NewInvalidEmptyParamError("customer_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return application.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
