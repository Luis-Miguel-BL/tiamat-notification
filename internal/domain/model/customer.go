package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

var AggregateTypeCustomer = domain.AggregateType("customer")

type CustomerID string

type SerializedCustomer struct {
	Attributes vo.CustomAttributes
	Events     map[vo.Slug]vo.CustomAttributes
}

type Customer struct {
	*domain.AggregateRoot
	customerID       CustomerID
	workspaceID      WorkspaceID
	externalID       vo.ExternalID
	name             vo.PersonName
	contact          vo.Contact
	customAttributes vo.CustomAttributes
	events           map[vo.Slug][]CustomerEvent
	segments         map[SegmentID]CustomerSegment
	createdAt        time.Time
	updatedAt        time.Time
}

type NewCustomerInput struct {
	ExternalID       vo.ExternalID
	WorkspaceID      WorkspaceID
	Name             vo.PersonName
	Contact          vo.Contact
	CustomAttributes vo.CustomAttributes
}

func NewCustomer(input NewCustomerInput) (customer *Customer, err domain.DomainError) {
	customerID := CustomerID(util.NewUUID())
	if input.ExternalID == "" {
		input.ExternalID, _ = vo.NewExternalID(util.NewUUID())
	}
	if input.WorkspaceID == "" {
		return customer, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	customer = &Customer{
		AggregateRoot:    domain.NewAggregateRoot(AggregateTypeCustomer, domain.AggregateID(customerID)),
		customerID:       customerID,
		externalID:       input.ExternalID,
		workspaceID:      input.WorkspaceID,
		name:             input.Name,
		contact:          input.Contact,
		customAttributes: input.CustomAttributes,
		events:           make(map[vo.Slug][]CustomerEvent),
		segments:         make(map[SegmentID]CustomerSegment),
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
	}

	customer.AggregateRoot.AppendEvent(event.CustomerCreatedEvent{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerCreatedEventType,
			OccurredAt:    customer.createdAt,
			AggregateType: customer.AggregateType(),
			AggregateID:   customer.AggregateID(),
		}),
		CustomerID:  string(customer.customerID),
		WorkspaceID: string(customer.workspaceID),
		CreatedAt:   customer.createdAt,
	})

	return customer, nil
}

type RestoreToRepoInput struct {
	CustomerID       CustomerID
	ExternalID       vo.ExternalID
	WorkspaceID      WorkspaceID
	Name             vo.PersonName
	Contact          vo.Contact
	CustomAttributes vo.CustomAttributes
	Events           map[vo.Slug][]CustomerEvent
	Segments         map[SegmentID]CustomerSegment
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func RestoreToRepo(input RestoreToRepoInput) (customer *Customer, err domain.DomainError) {
	if input.CustomerID == "" {
		return customer, domain.NewInvalidEmptyParamError("CustomerID")
	}
	if input.ExternalID == "" {
		return customer, domain.NewInvalidEmptyParamError("ExternalID")
	}
	if input.WorkspaceID == "" {
		return customer, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	if len(input.Events) == 0 {
		input.Events = make(map[vo.Slug][]CustomerEvent)
	}
	if len(input.Segments) == 0 {
		input.Segments = make(map[SegmentID]CustomerSegment)
	}
	customer = &Customer{
		AggregateRoot:    domain.NewAggregateRoot(AggregateTypeCustomer, domain.AggregateID(input.CustomerID)),
		customerID:       input.CustomerID,
		externalID:       input.ExternalID,
		workspaceID:      input.WorkspaceID,
		name:             input.Name,
		contact:          input.Contact,
		customAttributes: input.CustomAttributes,
		events:           input.Events,
		segments:         input.Segments,
		createdAt:        input.CreatedAt,
		updatedAt:        input.UpdatedAt,
	}
	return customer, nil
}

func (e *Customer) CustomerID() CustomerID {
	return e.customerID
}

func (e *Customer) WorkspaceID() WorkspaceID {
	return e.workspaceID
}
func (e *Customer) ExternalID() vo.ExternalID {
	return e.externalID
}
func (e *Customer) Name() vo.PersonName {
	return e.name
}

type UpdateCustomerInput struct {
	Name       vo.PersonName
	Email      vo.EmailAddress
	Phone      vo.PhoneNumber
	CustomAttr vo.CustomAttributes
}

func (e *Customer) Update(input UpdateCustomerInput) {
	if util.IsEmpty(input.Name.String()) {
		e.name = input.Name
	}
	if util.IsEmpty(input.Email.String()) {
		e.contact.Email.EmailAddress = input.Email
	}
	if util.IsEmpty(input.Phone.String()) {
		e.contact.Phone.PhoneNumber = input.Phone
	}
	for attrKey, attrValue := range input.CustomAttr {
		e.customAttributes[attrKey] = attrValue
	}

	e.updatedAt = time.Now()

	e.AggregateRoot.AppendEvent(event.CustomerUpdatedEvent{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerUpdatedEventType,
			OccurredAt:    e.createdAt,
			AggregateType: e.AggregateType(),
			AggregateID:   e.AggregateID(),
		}),
		CustomerID:  string(e.customerID),
		WorkspaceID: string(e.workspaceID),
		UpdatedAt:   e.updatedAt,
	})
}

func (e *Customer) CustomAttributes() vo.CustomAttributes {
	return e.customAttributes
}
func (e *Customer) Contact() vo.Contact {
	return e.contact
}
func (e *Customer) Events() map[vo.Slug][]CustomerEvent {
	return e.events
}
func (e *Customer) Segments() map[SegmentID]CustomerSegment {
	return e.segments
}
func (e *Customer) CreatedAt() time.Time {
	return e.createdAt
}
func (e *Customer) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *Customer) Serialize() (serialized SerializedCustomer) {
	serialized.Attributes = e.customAttributes.Flatten()
	serialized.Attributes["name"] = e.name.String()
	serialized.Attributes["first_name"] = e.name.GetFirstName()
	serialized.Attributes["email"] = e.contact.Email.EmailAddress.String()
	serialized.Attributes["phone"] = e.contact.Phone.PhoneNumber.String()

	serialized.Events = make(map[vo.Slug]vo.CustomAttributes)
	for eventSlug := range e.events {
		lastEvent := e.GetLastOccurrenceOfEvent(eventSlug)

		serialized.Events[eventSlug] = lastEvent.customAttributes.Flatten()
		serialized.Events[eventSlug]["occurred_at"] = lastEvent.OccurredAt()
	}
	return serialized
}

func (e *Customer) GetLastOccurrenceOfEvent(eventSlug vo.Slug) (lastEvent CustomerEvent) {
	for _, event := range e.events[eventSlug] {
		if event.OccurredAt().After(lastEvent.OccurredAt()) {
			lastEvent = event
		}
	}
	return lastEvent
}

func (e *Customer) AppendCustomerSegment(satisfiedSegment CustomerSegment) {
	e.segments[satisfiedSegment.SegmentID()] = satisfiedSegment
}

func (e *Customer) AppendCustomerEvent(slug vo.Slug, customAttributes vo.CustomAttributes) (err error) {
	customerEvent, err := NewCustomerEvent(NewCustomerEventInput{
		CustomerID:       e.customerID,
		WorkspaceID:      e.workspaceID,
		Slug:             slug,
		CustomAttributes: customAttributes,
	})
	if err != nil {
		return err
	}
	e.events[slug] = append(e.events[slug], *customerEvent)

	e.AggregateRoot.AppendEvent(
		event.CustomerEventOccurredEvent{
			DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
				EventType:     event.CustomerEventOccurredEventType,
				OccurredAt:    customerEvent.OccurredAt(),
				AggregateType: e.AggregateType(),
				AggregateID:   e.AggregateID(),
			}),
			CustomerID:      string(e.customerID),
			WorkspaceID:     string(e.workspaceID),
			CustomerEventID: string(customerEvent.customerEventID),
		})
	return nil
}
