package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type SaveCustomerInput struct {
	WorkspaceID      string
	ExternalID       string
	Name             string
	Contact          Contact
	CustomAttributes map[string]any
}

type Contact struct {
	EmailAddress string
	PhoneNumber  string
}

func (c *SaveCustomerInput) Validate() (err error) {
	if util.IsEmpty(c.WorkspaceID) {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Name) {
		return errors.NewInvalidEmptyParamError("name")
	}
	if util.IsEmpty(c.ExternalID) {
		return errors.NewInvalidEmptyParamError("external_id")
	}
	if util.IsEmpty(c.Contact.EmailAddress) && util.IsEmpty(c.Contact.PhoneNumber) {
		return errors.NewInvalidEmptyParamError("contact")
	}
	return nil
}
