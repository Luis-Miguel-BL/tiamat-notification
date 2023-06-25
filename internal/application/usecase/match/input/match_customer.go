package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type MatchCustomerInput struct {
	CustomerID  string
	WorkspaceID string
}

func (c *MatchCustomerInput) Validate() (err error) {
	if util.IsEmpty(c.CustomerID) {
		return errors.NewInvalidEmptyParamError("customer_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
