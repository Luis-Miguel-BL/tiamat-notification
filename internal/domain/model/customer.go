package model

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
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
	name             vo.PersonName
	contact          vo.Contact
	customAttributes vo.CustomAttributes
	events           map[vo.Slug][]CustomerEvent
	journeys         map[CampaignID]map[ActionID]CustomerJourney
	segments         map[SegmentID]CustomerSegment
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
		journeys:         make(map[CampaignID]map[ActionID]CustomerJourney),
		segments:         make(map[SegmentID]CustomerSegment),
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

func (e *Customer) GetJourney(campaignID CampaignID, actionID ActionID) (customerJourney CustomerJourney, found bool) {
	journey, found := e.journeys[campaignID][actionID]

	return journey, found
}
func (e *Customer) GetCampaignJourney(campaignID CampaignID) (customerJourney CustomerJourney, found bool) {
	actionsTriggered, found := e.journeys[campaignID]
	if !found {
		return customerJourney, found
	}
	for _, action := range actionsTriggered { // get last action triggered
		if action.TriggeredAt().After(customerJourney.TriggeredAt()) {
			customerJourney = action
		}
	}
	return customerJourney, found
}

func (e *Customer) AppendCustomerSegment(satisfiedSegment CustomerSegment) {
	e.segments[satisfiedSegment.SegmentID()] = satisfiedSegment
}

func (e *Customer) AppendJorney(customerJourney CustomerJourney) (err error) {
	e.journeys[customerJourney.campaignID][customerJourney.actionID] = customerJourney
	return nil
}

// func (e *Customer) GetActionTriggered(campaignID CampaignID, actionID ActionID) (actionTriggered ActionTriggered, found bool) {
// 	actionTriggered, found = e.actionsTriggered[campaignID][actionID]
// 	return actionTriggered, found
// }
