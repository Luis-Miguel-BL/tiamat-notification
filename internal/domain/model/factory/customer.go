package factory

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type CustomerFactory struct{}

type CreateCustomerInput struct {
	ExternalID       string
	WorkspaceID      string
	Name             string
	CustomAttributes map[string]any
}

func (f *CustomerFactory) CreateCustomer(input CreateCustomerInput) (customer *model.Customer) {
	if input.ExternalID == "" {
		input.ExternalID = "fake-external-id"
	}
	if input.WorkspaceID == "" {
		input.WorkspaceID = "fake-workspace-id"
	}
	if input.Name == "" {
		input.Name = "fake-name"
	}
	if input.CustomAttributes == nil {
		input.CustomAttributes = make(map[string]any)
	}

	customer, _ = model.NewCustomer(model.NewCustomerInput{
		ExternalID:       vo.ExternalID(input.ExternalID),
		WorkspaceID:      model.WorkspaceID(input.WorkspaceID),
		Name:             vo.PersonName(input.Name),
		CustomAttributes: input.CustomAttributes,
	})

	return customer
}
