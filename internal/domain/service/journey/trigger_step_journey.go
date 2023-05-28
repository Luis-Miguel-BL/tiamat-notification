package journey

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type TriggerStepJourneyService struct {
	gatewayManager gateway.GatewayManager
}

func NewTriggerStepJourneyService(gatewayManager gateway.GatewayManager) TriggerStepJourneyService {
	return TriggerStepJourneyService{
		gatewayManager: gatewayManager,
	}
}

func (s *TriggerStepJourneyService) TriggerStepJourney(ctx context.Context, customer *model.Customer, action model.Action, campaign model.Campaign) (err error) {
	stepJourney, found := customer.GetStepJourney(campaign.CampaignID(), action.ActionID())
	if !found {
		return domain.NewNotFoundError("customer-journey")
	}
	eventBase := domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
		EventType:     event.CustomerEventOccurredEventType,
		OccurredAt:    stepJourney.TriggeredAt(),
		AggregateType: customer.AggregateType(),
		AggregateID:   customer.AggregateID(),
	})

	matchFilters := matchFilters(customer, campaign.Filters())
	if matchFilters {
		stepJourney.Finish(model.StepJourneySkipped, "")

		customer.AggregateRoot.AppendEvent(event.StepJourneySkipped{
			DomainEventBase: eventBase,
			CustomerID:      string(customer.CustomerID()),
			WorkspaceID:     string(customer.WorkspaceID()),
			CampaignID:      string(stepJourney.CampaignID()),
			ActionID:        string(stepJourney.ActionID()),
			StepJourneyID:   string(stepJourney.StepJourneyID()),
			Reason:          "campaign-filter-matched",
			TriggeredAt:     stepJourney.TriggeredAt(),
		})
		return nil
	}

	if !action.IsActive() {
		stepJourney.Finish(model.StepJourneySkipped, "")

		customer.AggregateRoot.AppendEvent(event.StepJourneySkipped{
			DomainEventBase: eventBase,
			CustomerID:      string(customer.CustomerID()),
			WorkspaceID:     string(customer.WorkspaceID()),
			CampaignID:      string(stepJourney.CampaignID()),
			ActionID:        string(stepJourney.ActionID()),
			StepJourneyID:   string(stepJourney.StepJourneyID()),
			Reason:          "action-disabled",
			TriggeredAt:     stepJourney.TriggeredAt(),
		})
		return nil
	}

	status, nextActionID, err := mapActionHandler[action.BehaviorType()](ctx, s.gatewayManager, customer, action)

	stepJourney.Finish(status, nextActionID)

	switch status {
	case model.StepJourneyStatusSuccessed:
		customer.AggregateRoot.AppendEvent(event.StepJourneySuccessed{
			DomainEventBase: eventBase,
			CustomerID:      string(customer.CustomerID()),
			WorkspaceID:     string(customer.WorkspaceID()),
			CampaignID:      string(stepJourney.CampaignID()),
			ActionID:        string(stepJourney.ActionID()),
			StepJourneyID:   string(stepJourney.StepJourneyID()),
			TriggeredAt:     stepJourney.TriggeredAt(),
		})
	case model.StepJourneyStatusScheduled:
		customer.AggregateRoot.AppendEvent(event.StepJourneyScheduled{
			DomainEventBase: eventBase,
			CustomerID:      string(customer.CustomerID()),
			WorkspaceID:     string(customer.WorkspaceID()),
			CampaignID:      string(stepJourney.CampaignID()),
			ActionID:        string(stepJourney.ActionID()),
			StepJourneyID:   string(stepJourney.StepJourneyID()),
			TriggeredAt:     stepJourney.TriggeredAt(),
		})
	case model.StepJourneyStatusFailed:
		customer.AggregateRoot.AppendEvent(event.StepJourneyFailedType{
			DomainEventBase: eventBase,
			CustomerID:      string(customer.CustomerID()),
			WorkspaceID:     string(customer.WorkspaceID()),
			CampaignID:      string(stepJourney.CampaignID()),
			ActionID:        string(stepJourney.ActionID()),
			StepJourneyID:   string(stepJourney.StepJourneyID()),
			Description:     err.Error(),
			TriggeredAt:     stepJourney.TriggeredAt(),
		})
	default:
		return domain.NewInvalidParamError("JourneyStatus")
	}

	return nil
}

type Handle func(context.Context, gateway.GatewayManager, *model.Customer, model.Action) (model.StepJourneyStatus, model.ActionID, error)

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

func matchFilters(customer *model.Customer, campaignFilters []model.SegmentID) (isMatch bool) {
	customerSegments := customer.GetSegments()

	for _, segmentFilter := range campaignFilters {
		_, isMatch := customerSegments[segmentFilter]
		if !isMatch {
			return false
		}
	}
	return true
}

func handleIfAttribute(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {
	return
}
func handleRandom(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleSendEmail(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleSendSMS(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleSendWhatsapp(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleSplit(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleWaitFor(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {
	return
}

func handleWaitUntil(ctx context.Context, gatewayManager gateway.GatewayManager, customer *model.Customer, action model.Action) (status model.StepJourneyStatus, nextActionID model.ActionID, err error) {
	return
}
