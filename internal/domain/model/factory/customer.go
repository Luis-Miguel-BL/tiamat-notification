package factory

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type CustomerFactory struct{}

type CreateCustomerInput struct {
	CustomerID       string
	ExternalID       string
	WorkspaceID      string
	Name             string
	Contact          vo.Contact
	CustomAttributes map[string]any
	Events           []model.CustomerEvent
	Segments         []model.CustomerSegment
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (f *CustomerFactory) CreateCustomer(input CreateCustomerInput) (customer *model.Customer) {
	if input.CustomAttributes == nil {
		input.CustomAttributes = make(map[string]any)
	}

	customer, _ = model.NewCustomer(model.NewCustomerInput{
		Cus
		ExternalID:       vo.ExternalID(input.ExternalID),
		WorkspaceID:      model.WorkspaceID(input.WorkspaceID),
		Name:             vo.PersonName(input.Name),
		CustomAttributes: input.CustomAttributes,
	})

	return customer
}
