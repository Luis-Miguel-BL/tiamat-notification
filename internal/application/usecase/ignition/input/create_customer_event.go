package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CreateCustomerEventInput struct {
	CustomerID       string
	WorkspaceID      string
	CustomerEventID  string
	Slug             string
	CustomAttributes map[string]any
}

func (c *CreateCustomerEventInput) Validate() (err error) {
	if util.IsEmpty(c.CustomerEventID) {
		c.CustomerEventID = util.NewUUID()
	}
	if util.IsEmpty(c.CustomerID) {
		return errors.NewInvalidEmptyParamError("customer_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Slug) {
		return errors.NewInvalidEmptyParamError("slug")
	}
	return nil
}
