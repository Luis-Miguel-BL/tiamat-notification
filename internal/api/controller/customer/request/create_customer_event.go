package request

import "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"

type CreateCustomerEvent struct {
	CustomerID       string         `json:"customer_id"`
	WorkspaceID      string         `json:"workspace_id"`
	CustomerEventID  string         `json:"customer_event_id"`
	Slug             string         `json:"slug"`
	CustomAttributes map[string]any `json:"custom_attributes"`
}

func (r *CreateCustomerEvent) Validate() (err error) {
	if r.CustomerID == "" {
		return errors.NewInvalidEmptyParamError("customer_id")
	}
	if r.WorkspaceID == "" {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if r.CustomerEventID == "" {
		return errors.NewInvalidEmptyParamError("customer_event_id")
	}
	if r.Slug == "" {
		return errors.NewInvalidEmptyParamError("slug")
	}
	return nil
}
