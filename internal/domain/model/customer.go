package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

var AggregateTypeCustomer = domain.AggregateType("customer")

type CustomerID string

func NewCustomerID(customerID string) CustomerID {
	return CustomerID(customerID)
}

type SerializedCustomer struct {
	Attributes vo.CustomAttributes
	Events     map[vo.Slug]vo.CustomAttributes
}

type Customer struct {
	*domain.AggregateRoot
	customerID       CustomerID
	workspaceID      WorkspaceID
	name             vo.PersonName
	contact          vo.Contact
	customAttributes vo.CustomAttributes
	events           map[vo.Slug][]CustomerEvent
	actionsTriggered map[CampaignID][]ActionTriggered
	currentSegments  map[SegmentID]CurrentSegment
	createdAt        time.Time
	updatedAt        time.Time
}

type NewCustomerInput struct {
	CustomerID       CustomerID
	WorkspaceID      WorkspaceID
	Name             vo.PersonName
	Contact          vo.Contact
	CustomAttributes vo.CustomAttributes
}

func NewCustomer(input NewCustomerInput) (customer *Customer, err domain.DomainError) {
	if input.CustomerID == "" {
		return customer, domain.NewInvalidEmptyParamError("CustomerID")
	}
	if input.WorkspaceID == "" {
		return customer, domain.NewInvalidEmptyParamError("WorkspaceID")
	}
	customer = &Customer{
		AggregateRoot:    domain.NewAggregateRoot(AggregateTypeCustomer, domain.AggregateID(input.CustomerID)),
		customerID:       input.CustomerID,
		workspaceID:      input.WorkspaceID,
		name:             input.Name,
		contact:          input.Contact,
		customAttributes: input.CustomAttributes,
		events:           make(map[vo.Slug][]CustomerEvent),
		actionsTriggered: make(map[CampaignID][]ActionTriggered),
		currentSegments:  make(map[SegmentID]CurrentSegment),
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
	}

	customer.AggregateRoot.AppendEvent(event.CustomerCreatedEvent{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerEventOccurredEventType,
			OccurredAt:    customer.createdAt,
			AggregateType: customer.AggregateType(),
			AggregateID:   customer.AggregateID(),
		}),
		CustomerID:       string(customer.customerID),
		WorkspaceID:      string(customer.workspaceID),
		Name:             customer.name,
		Contact:          customer.contact,
		CustomAttributes: customer.customAttributes,
		CreatedAt:        customer.createdAt,
	})

	return customer, nil
}

func (e *Customer) CustomerID() CustomerID {
	return e.customerID
}

func (e *Customer) WorkspaceID() WorkspaceID {
	return e.workspaceID
}

func (e *Customer) Serialize() (serialized SerializedCustomer) {
	serialized.Attributes = e.customAttributes
	serialized.Attributes["name"] = e.name
	serialized.Attributes["first_name"] = e.name.GetFirstName()
	serialized.Attributes["email"] = e.contact.Email.EmailAddress
	serialized.Attributes["phone"] = e.contact.Phone.PhoneNumber

	serialized.Events = make(map[vo.Slug]vo.CustomAttributes)
	for eventSlug, _ := range e.events {
		lastEvent := e.GetLastOccurrenceOfEvent(eventSlug)

		serialized.Events[eventSlug] = lastEvent.customAttributes
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

func (e *Customer) GetLastActionTrigged(campaignID CampaignID) (lastActionTrigged ActionTriggered, alreadyTriggered bool) {
	actionsTriggered, alreadyTriggered := e.actionsTriggered[campaignID]
	if !alreadyTriggered {
		return lastActionTrigged, alreadyTriggered
	}
	for _, action := range actionsTriggered {
		if action.TriggeredAt().After(lastActionTrigged.TriggeredAt()) {
			lastActionTrigged = action
		}
	}
	return lastActionTrigged, alreadyTriggered
}

func (e *Customer) AppendCurrentSegment(currentSegment CurrentSegment) {
	e.currentSegments[currentSegment.SegmentID()] = currentSegment

	e.AggregateRoot.AppendEvent(event.CustomerMatched{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerEventOccurredEventType,
			OccurredAt:    currentSegment.MatchedAt(),
			AggregateType: e.AggregateType(),
			AggregateID:   e.AggregateID(),
		}),
		CustomerID:  string(e.customerID),
		WorkspaceID: string(e.workspaceID),
		SegmentID:   string(currentSegment.SegmentID()),
		MatchedAt:   currentSegment.MatchedAt(),
	})
}
