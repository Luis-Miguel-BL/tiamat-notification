package service

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type TriggerService interface {
	TriggerAction(ctx context.Context, customer *model.Customer, action model.Action, campaignFilters []model.Segment, campaignID model.CampaignID) (err error)
}

type triggerService struct {
	gatewayManager gateway.GatewayManager
}

func NewTriggerService(gatewayManager gateway.GatewayManager) TriggerService {
	return &triggerService{
		gatewayManager: gatewayManager,
	}
}

type Handle func(context.Context, gateway.GatewayManager, *model.Customer, model.Action) (model.CustomerJourneyStatus, model.ActionID, error)

var mapActionHandler = map[model.BehaviorType]Handle{
	model.BehaviorTypeSendEmail:    handleSendEmail,
	model.BehaviorTypeSendSMS:      handleSendSMS,
	model.BehaviorTypeSendWhatsapp: handleSendWhatsapp,
	model.BehaviorTypeWaitFor:      handleWaitFor,
	model.BehaviorTypeWaitUntil:    handleWaitUntil,
	model.BehaviorTypeIfAttribute:  handleIfAttribute,
	model.BehaviorTypeRandom:       handleRandom,
	model.BehaviorTypeSplit:        handleSplit,
}

func (s *triggerService) TriggerAction(ctx context.Context, customer *model.Customer, action model.Action, campaignFilters []model.Segment, campaignID model.CampaignID) (err error) {
	customerJourney, found := customer.GetJourney(campaignID, action.ActionID())
	if !found {
		return domain.NewNotFoundError("customer-journey")
	}
	matchFilters, err := matchFilters(campaignFilters, customer)
	if err != nil {
		return err
	}
	if matchFilters {
		customerJourney.Finish(model.CustomerJourneyStatusFilterMatch, "")
	}
	status, nextActionID, err := mapActionHandler[action.BehaviorType()](ctx, s.gatewayManager, customer, action)
	if err != nil {
		return err
	}

	customerJourney.Finish(status, nextActionID)

	eventBase := domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
		EventType:     event.CustomerEventOccurredEventType,
		OccurredAt:    customerJourney.TriggeredAt(),
		AggregateType: customer.AggregateType(),
		AggregateID:   customer.AggregateID(),
	})

	switch status {
	case model.CustomerJourneyStatusSuccess:
	case model.CustomerJourneyStatusScheduled:
	case model.CustomerJourneyStatusFailed:
	case model.CustomerJourneyStatusFilterMatch:
		customer.AggregateRoot.AppendEvent(event.CampaignFilterMatched{
			DomainEventBase:   eventBase,
			CustomerID:        string(customer.CustomerID()),
			WorkspaceID:       string(customer.WorkspaceID()),
			CampaignID:        string(customerJourney.CampaignID()),
			ActionID:          string(customerJourney.ActionID()),
			CustomerJourneyID: string(customerJourney.CustomerJourneyID()),
			TriggeredAt:       customerJourney.TriggeredAt(),
		})
	default:
		return domain.NewInvalidParamError("JourneyStatus")
	}

	return nil
}

func matchFilters(campaignFilters []model.Segment, customer *model.Customer) (isMatchAll bool, err error) {
	isMatchAll = true
	for _, segment := range campaignFilters {
		for _, condition := range segment.Conditions() {
			if !condition.IsMatch(customer.Serialize()) {
				return false, nil
			}
		}

		satisfiedSegment, err := model.NewCustomerSegment(
			model.NewCustomerSegmentInput{
				CustomerID:  customer.CustomerID(),
				WorkspaceID: customer.WorkspaceID(),
				SegmentID:   segment.SegmentID(),
			},
		)
		if err != nil {
			return false, err
		}
		customer.AppendCustomerSegment(*satisfiedSegment)
	}
	return isMatchAll, nil
}

func handleIfAttribute(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.CustomerJourneyStatus, nextActionID model.ActionID, err error) {
	return
}
func handleRandom(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.CustomerJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleSendEmail(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.CustomerJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleSendSMS(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.CustomerJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleSendWhatsapp(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.CustomerJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleSplit(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.CustomerJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleWaitFor(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.CustomerJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleWaitUntil(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.CustomerJourneyStatus, nextActionID model.ActionID, err error) {
	return
}
