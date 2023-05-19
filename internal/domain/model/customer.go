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
	customerID        CustomerID
	workspaceID       WorkspaceID
	name              vo.PersonName
	contact           vo.Contact
	customAttributes  vo.CustomAttributes
	events            map[vo.Slug][]CustomerEvent
	actionsTriggered  map[CampaignID]map[ActionID]ActionTriggered
	satisfiedSegments map[SegmentID]SatisfiedSegment
	createdAt         time.Time
	updatedAt         time.Time
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
		AggregateRoot:     domain.NewAggregateRoot(AggregateTypeCustomer, domain.AggregateID(input.CustomerID)),
		customerID:        input.CustomerID,
		workspaceID:       input.WorkspaceID,
		name:              input.Name,
		contact:           input.Contact,
		customAttributes:  input.CustomAttributes,
		events:            make(map[vo.Slug][]CustomerEvent),
		actionsTriggered:  make(map[CampaignID]map[ActionID]ActionTriggered),
		satisfiedSegments: make(map[SegmentID]SatisfiedSegment),
		createdAt:         time.Now(),
		updatedAt:         time.Now(),
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

func (e *Customer) AppendSatisfiedSegment(satisfiedSegment SatisfiedSegment) {
	e.satisfiedSegments[satisfiedSegment.SegmentID()] = satisfiedSegment

	e.AggregateRoot.AppendEvent(event.CustomerMatched{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerEventOccurredEventType,
			OccurredAt:    satisfiedSegment.MatchedAt(),
			AggregateType: e.AggregateType(),
			AggregateID:   e.AggregateID(),
		}),
		CustomerID:  string(e.customerID),
		WorkspaceID: string(e.workspaceID),
		SegmentID:   string(satisfiedSegment.SegmentID()),
		MatchedAt:   satisfiedSegment.MatchedAt(),
	})
}

func (e *Customer) TriggerAction(action Action) (err error) {
	actionTriggered, err := NewActionTriggered(NewActionTriggeredInput{
		WorkspaceID: e.workspaceID,
		CustomerID:  e.customerID,
		CampaignID:  action.campaignID,
		ActionID:    action.actionID,
	})
	if err != nil {
		return err
	}

	e.actionsTriggered[action.campaignID][action.actionID] = *actionTriggered

	e.AggregateRoot.AppendEvent(event.ActionTrigged{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerEventOccurredEventType,
			OccurredAt:    actionTriggered.TriggeredAt(),
			AggregateType: e.AggregateType(),
			AggregateID:   e.AggregateID(),
		}),
		CustomerID:        string(e.customerID),
		WorkspaceID:       string(e.workspaceID),
		CampaignID:        string(actionTriggered.campaignID),
		ActionID:          string(actionTriggered.actionID),
		ActionTriggeredID: string(actionTriggered.actionTriggeredID),
		TriggeredAt:       actionTriggered.triggeredAt,
	})
	return nil
}

// func (e *Customer) GetActionTriggered(campaignID CampaignID, actionID ActionID) (actionTriggered ActionTriggered, found bool) {
// 	actionTriggered, found = e.actionsTriggered[campaignID][actionID]
// 	return actionTriggered, found
// }

func (e *Customer) FinishActionTriggered(action Action, status ActionTriggeredStatus, nextActionID ActionID) (err error) {
	actionTriggered, ok := e.actionsTriggered[action.campaignID][action.actionID]
	if !ok {
		return domain.NewNotFoundError("action-triggered")
	}
	actionTriggered.status = status
	e.actionsTriggered[action.campaignID][action.actionID] = actionTriggered

	eventBase := domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
		EventType:     event.CustomerEventOccurredEventType,
		OccurredAt:    actionTriggered.TriggeredAt(),
		AggregateType: e.AggregateType(),
		AggregateID:   e.AggregateID(),
	})

	switch status {
	case ActionTriggeredStatusSuccess:
	case ActionTriggeredStatusScheduled:
	case ActionTriggeredStatusFailed:
	case ActionTriggeredStatusFilterMatch:
		e.AggregateRoot.AppendEvent(event.CampaignFilterMatched{
			DomainEventBase:   eventBase,
			CustomerID:        string(e.customerID),
			WorkspaceID:       string(e.workspaceID),
			CampaignID:        string(actionTriggered.campaignID),
			ActionID:          string(actionTriggered.actionID),
			ActionTriggeredID: string(actionTriggered.actionTriggeredID),
			TriggeredAt:       actionTriggered.triggeredAt,
		})
	default:
		return domain.NewInvalidParamError("ActionTriggeredStatus")
	}

	return nil
}
