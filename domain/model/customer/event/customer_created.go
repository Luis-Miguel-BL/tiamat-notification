package event

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain"
	entity "github.com/Luis-Miguel-BL/tiamat-notification/domain/model/customer"
)

var CustomerCreatedEventType = domain.EventType("customer-created")

type CustomerCreatedEvent struct {
	*domain.DomainEventBase
	CustomerCreated entity.Customer
}

func NewCustomerCreatedEvent(customerCreated entity.Customer) (event *CustomerCreatedEvent) {
	event = &CustomerCreatedEvent{
		DomainEventBase: &domain.DomainEventBase{
			EventType:     CustomerCreatedEventType,
			OccurredAt:    customerCreated.CreatedAt,
			AggregateType: entity.AggregateType,
			AggregateID:   customerCreated.AggregateID,
		},
		CustomerCreated: customerCreated,
	}
	return event
}

func (e CustomerCreatedEvent) EventType() domain.EventType {
	return e.DomainEventBase.EventType
}
func (e CustomerCreatedEvent) OccurredAt() time.Time {
	return e.DomainEventBase.OccurredAt
}
func (e CustomerCreatedEvent) AggregateType() domain.AggregateType {
	return e.DomainEventBase.AggregateType
}
func (e CustomerCreatedEvent) AggregateID() domain.AggregateID {
	return e.DomainEventBase.AggregateID
}
