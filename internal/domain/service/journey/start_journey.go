package journey

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
)

type StartJourneyService struct {
}

func NewStartJourneyService(repo repository.CustomerRepository) StartJourneyService {
	return StartJourneyService{}
}

func (s *StartJourneyService) StartJourney(ctx context.Context, customer *model.Customer, campaign model.Campaign) (err error) {
	lastTriggered, found := customer.GetCampaignJourney(campaign.CampaignID())
	if found {
		if !campaign.MustBeTriggered(lastTriggered.TriggeredAt()) {
			return nil
		}
	}
	firstAction, err := campaign.Action(campaign.FirstActionID())
	if err != nil {
		return err
	}
	stepJourney, err := model.NewStepJourney(model.NewStepJourneyInput{
		WorkspaceID: customer.WorkspaceID(),
		CustomerID:  customer.CustomerID(),
		CampaignID:  campaign.CampaignID(),
		ActionID:    firstAction.ActionID(),
	})
	if err != nil {
		return err
	}
	err = customer.AppendJourney(*stepJourney)
	if err != nil {
		return err
	}

	customer.AggregateRoot.AppendEvent(event.ActionTrigged{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerEventOccurredEventType,
			OccurredAt:    stepJourney.TriggeredAt(),
			AggregateType: customer.AggregateType(),
			AggregateID:   customer.AggregateID(),
		}),
		CustomerID:    string(customer.CustomerID()),
		WorkspaceID:   string(customer.WorkspaceID()),
		CampaignID:    string(campaign.CampaignID()),
		ActionID:      string(firstAction.ActionID()),
		StepJourneyID: string(stepJourney.StepJourneyID()),
		TriggeredAt:   stepJourney.TriggeredAt(),
	})

	return nil
}
