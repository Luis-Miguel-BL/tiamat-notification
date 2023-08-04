package request

import "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"

type SaveCustomer struct {
	WorkspaceID      string         `json:"workspace_id"`
	ExternalID       string         `json:"external_id"`
	Name             string         `json:"name"`
	EmailAddress     string         `json:"email_address"`
	PhoneNumber      string         `json:"phone_number"`
	CustomAttributes map[string]any `json:"custom_attributes"`
}

func (r *SaveCustomer) Validate() (err error) {
	if r.WorkspaceID == "" {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if r.ExternalID == "" {
		return errors.NewInvalidEmptyParamError("external_id")
	}
	if r.Name == "" {
		return errors.NewInvalidEmptyParamError("name")
	}
	if r.EmailAddress == "" && r.PhoneNumber == "" {
		return errors.NewInvalidEmptyParamError("email_address")
	}
	return nil
}
