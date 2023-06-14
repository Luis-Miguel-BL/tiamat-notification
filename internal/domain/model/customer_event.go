package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type CustomerEventID string

func NewCustomerEventID(customerEventID string) CustomerEventID {
	return CustomerEventID(customerEventID)
}

type CustomerEvent struct {
	customerEventID  CustomerEventID
	customerID       CustomerID
	workspaceID      WorkspaceID
	slug             vo.Slug
	customAttributes vo.CustomAttributes
	occurredAt       time.Time
}

type NewCustomerEventInput struct {
	CustomerID       CustomerID
	WorkspaceID      WorkspaceID
	Slug             vo.Slug
	CustomAttributes vo.CustomAttributes
}

func NewCustomerEvent(input NewCustomerEventInput) (customerEvent *CustomerEvent, err domain.DomainError) {
	if input.Slug == "" {
		return customerEvent, domain.NewInvalidEmptyParamError("Slug")
	}
	customerEvent = &CustomerEvent{
		customerEventID:  CustomerEventID(util.NewUUID()),
		customerID:       input.CustomerID,
		workspaceID:      input.WorkspaceID,
		slug:             input.Slug,
		customAttributes: input.CustomAttributes,
		occurredAt:       time.Now(),
	}

	return customerEvent, nil
}

func (e *CustomerEvent) OccurredAt() time.Time {
	return e.occurredAt
}
func (e *CustomerEvent) CustomerEventID() CustomerEventID {
	return e.customerEventID
}
func (e *CustomerEvent) Slug() vo.Slug {
	return e.slug
}
func (e *CustomerEvent) CustomAttributes() vo.CustomAttributes {
	return e.customAttributes
}
