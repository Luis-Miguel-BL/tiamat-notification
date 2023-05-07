package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type CustomerEventID string

func NewCustomerEventID(customerEventID string) CustomerEventID {
	return CustomerEventID(customerEventID)
}

type CustomerEvent struct {
	*domain.AggregateRoot
	customerEventID  CustomerEventID
	customerID       CustomerID
	workspaceID      WorkspaceID
	slug             vo.Slug
	customAttributes vo.CustomAttributes
	occurredAt       time.Time
}

type NewCustomerEventInput struct {
	CustomerEventID  CustomerEventID
	CustomerID       CustomerID
	WorkspaceID      WorkspaceID
	Slug             vo.Slug
	CustomAttributes vo.CustomAttributes
}

func NewCustomerEvent(input NewCustomerEventInput) (customerEvent *CustomerEvent, err domain.DomainError) {
	if input.CustomerEventID == "" {
		return customerEvent, domain.NewInvalidEmptyParamError("CustomerEventID")
	}
	if input.Slug == "" {
		return customerEvent, domain.NewInvalidEmptyParamError("Slug")
	}
	customerEvent = &CustomerEvent{
		AggregateRoot:    domain.NewAggregateRoot(AggregateTypeCustomer, domain.AggregateID(input.CustomerID)),
		customerEventID:  input.CustomerEventID,
		customerID:       input.CustomerID,
		workspaceID:      input.WorkspaceID,
		slug:             input.Slug,
		customAttributes: input.CustomAttributes,
		occurredAt:       time.Now(),
	}
	customerEvent.AggregateRoot.AppendEvent(
		event.CustomerEventOccurredEvent{
			DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
				EventType:     event.CustomerEventOccurredEventType,
				OccurredAt:    customerEvent.OccurredAt(),
				AggregateType: customerEvent.AggregateType(),
				AggregateID:   customerEvent.AggregateID(),
			}),
			CustomerID:       string(input.CustomerID),
			WorkspaceID:      string(input.WorkspaceID),
			CustomerEventID:  string(input.CustomerEventID),
			Slug:             customerEvent.Slug(),
			CustomAttributes: customerEvent.CustomAttributes(),
		})

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
