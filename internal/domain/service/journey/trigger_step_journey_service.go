package journey

import (
	"context"
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/journey/handle_action"
)

type TriggerStepJourneyService struct {
	gatewayManager gateway.GatewayManager
}

func NewTriggerStepJourneyService(gatewayManager gateway.GatewayManager) TriggerStepJourneyService {
	return TriggerStepJourneyService{
		gatewayManager: gatewayManager,
	}
}

func (s *TriggerStepJourneyService) TriggerStepJourney(ctx context.Context, journey *model.Journey, customer model.Customer, action model.Action, campaign model.Campaign) (err error) {
	stepJourney, found := journey.StepJourney(action.ActionID())
	if !found {
		return domain.NewNotFoundError("journey-step")
	}

	matchFilters := matchFilters(customer, campaign.Filters())
	if matchFilters {
		journey.SkippStep(stepJourney, event.SkippedReasonMatchFilters)
		return nil
	}

	if !action.IsActive() {
		journey.SkippStep(stepJourney, event.SkippedReasonActionDisabled)
		return nil
	}

	isNotificationAction := action.ActionType() == model.ActionTypeNotification
	isAvailableToSendNotification, nextAvailableTime := campaign.NextAvailableTimeToTriggerNotification(time.Now())
	if isNotificationAction && !isAvailableToSendNotification {
		s.gatewayManager.SchedulerGateway.Scheduler(nextAvailableTime)

		journey.AggregateRoot.AppendEvent(event.StepJourneyScheduled{
			DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
				EventType:     event.StepJourneyScheduledEventType,
				OccurredAt:    time.Now(),
				AggregateType: model.AggregateTypeJourney,
				AggregateID:   journey.AggregateID(),
			}),
			CustomerID:    string(journey.CustomerID()),
			WorkspaceID:   string(journey.WorkspaceID()),
			CampaignID:    string(journey.CampaignID()),
			ActionID:      string(stepJourney.ActionID()),
			StepJourneyID: string(stepJourney.StepJourneyID()),
			JourneyID:     string(journey.JourneyID()),
			Reason:        event.ScheduledReasonOutOfNotificationTimeRange,
			TriggeredAt:   stepJourney.TriggeredAt(),
		})

	}

	status, nextActionID, err := handle_action.Handle(ctx, s.gatewayManager, action, customer)

	switch status {
	case model.StepJourneyStatusSuccessed:
		journey.FinishSuccessfully(stepJourney, nextActionID)
	case model.StepJourneyStatusFailed:
		journey.FinishUnsuccessfully(stepJourney, nextActionID, err.Error())
	case model.StepJourneyStatusScheduled:
		journey.AggregateRoot.AppendEvent(event.StepJourneyScheduled{
			DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
				EventType:     event.StepJourneyScheduledEventType,
				OccurredAt:    time.Now(),
				AggregateType: model.AggregateTypeJourney,
				AggregateID:   journey.AggregateID(),
			}),
			CustomerID:    string(journey.CustomerID()),
			WorkspaceID:   string(journey.WorkspaceID()),
			CampaignID:    string(journey.CampaignID()),
			ActionID:      string(stepJourney.ActionID()),
			StepJourneyID: string(stepJourney.StepJourneyID()),
			JourneyID:     string(journey.JourneyID()),
			Reason:        event.ScheduledReasonScheduledByAction,
			TriggeredAt:   stepJourney.TriggeredAt(),
		})

	default:
		return domain.NewInvalidParamError("JourneyStatus")
	}

	return nil
}

func matchFilters(customer model.Customer, campaignFilters []model.SegmentID) (isMatch bool) {
	customerSegments := customer.GetSegments()

	for _, segmentFilter := range campaignFilters {
		_, isMatch := customerSegments[segmentFilter]
		if !isMatch {
			return false
		}
	}
	return true
}
