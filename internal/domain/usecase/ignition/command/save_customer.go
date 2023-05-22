package command

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type SaveCustomerCommand struct {
	CustomerID       string         `json:"customer_id,omitempty"`
	WorkspaceID      string         `json:"workspace_id,omitempty"`
	Name             string         `json:"name,omitempty"`
	Contact          Contact        `json:"contact,omitempty"`
	CustomAttributes map[string]any `json:"custom_attributes,omitempty"`
}

type Contact struct {
	EmailAddress string `json:"email_address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
}

func (c *SaveCustomerCommand) Validate() (err error) {
	if util.IsEmpty(c.CustomerID) {
		c.CustomerID = util.NewUUID()
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if util.IsEmpty(c.Name) {
		return domain.NewInvalidEmptyParamError("name")
	}
	if util.IsEmpty(c.Contact.EmailAddress) && util.IsEmpty(c.Contact.PhoneNumber) {
		return domain.NewInvalidEmptyParamError("contact")
	}
	return nil
}
