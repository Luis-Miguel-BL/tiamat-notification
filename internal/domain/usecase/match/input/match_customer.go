package input

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type MatchCustomerInput struct {
	CustomerID  string `json:"customer_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}

func (c *MatchCustomerInput) Validate() (err error) {
	if util.IsEmpty(c.CustomerID) {
		return domain.NewInvalidEmptyParamError("customer_id")
	}
	if util.IsEmpty(c.WorkspaceID) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	return nil
}
