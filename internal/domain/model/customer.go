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
	journeys         map[CampaignID]map[ActionID]StepJourney
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
		journeys:         make(map[CampaignID]map[ActionID]StepJourney),
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
		ExternalID:       customer.externalID,
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

func (e *Customer) GetSegments() map[SegmentID]CustomerSegment {
	return e.segments
}

func (e *Customer) GetStepJourney(campaignID CampaignID, actionID ActionID) (stepJourney StepJourney, found bool) {
	journey, found := e.journeys[campaignID][actionID]

	return journey, found
}

func (e *Customer) GetCampaignJourney(campaignID CampaignID) (stepJourney StepJourney, found bool) {
	actionsTriggered, found := e.journeys[campaignID]
	if !found {
		return stepJourney, found
	}
	for _, action := range actionsTriggered { // get last action triggered
		if action.TriggeredAt().After(stepJourney.TriggeredAt()) {
			stepJourney = action
		}
	}
	return stepJourney, found
}

func (e *Customer) AppendCustomerSegment(satisfiedSegment CustomerSegment) {
	e.segments[satisfiedSegment.SegmentID()] = satisfiedSegment
}

func (e *Customer) AppendJourney(stepJourney StepJourney) (err error) {
	e.journeys[stepJourney.campaignID][stepJourney.actionID] = stepJourney
	return nil
}

// func (e *Customer) GetActionTriggered(campaignID CampaignID, actionID ActionID) (actionTriggered ActionTriggered, found bool) {
// 	actionTriggered, found = e.actionsTriggered[campaignID][actionID]
// 	return actionTriggered, found
// }
