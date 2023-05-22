package service

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
)

type JourneyService interface {
	StartJourney(ctx context.Context, customer *model.Customer, campaign model.Campaign) (err error)
}

type journeyService struct {
	repo repository.CustomerRepository
}

func NewJourneyService(repo repository.CustomerRepository) JourneyService {
	return &journeyService{
		repo: repo,
	}
}

func (s *journeyService) StartJourney(ctx context.Context, customer *model.Customer, campaign model.Campaign) (err error) {
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
	customerJourney, err := model.NewCustomerJourney(model.NewCustomerJourneyInput{
		WorkspaceID: customer.WorkspaceID(),
		CustomerID:  customer.CustomerID(),
		CampaignID:  campaign.CampaignID(),
		ActionID:    firstAction.ActionID(),
	})
	if err != nil {
		return err
	}
	err = customer.AppendJorney(*customerJourney)
	if err != nil {
		return err
	}

	customer.AggregateRoot.AppendEvent(event.ActionTrigged{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerEventOccurredEventType,
			OccurredAt:    customerJourney.TriggeredAt(),
			AggregateType: customer.AggregateType(),
			AggregateID:   customer.AggregateID(),
		}),
		CustomerID:        string(customer.CustomerID()),
		WorkspaceID:       string(customer.WorkspaceID()),
		CampaignID:        string(campaign.CampaignID()),
		ActionID:          string(firstAction.ActionID()),
		CustomerJourneyID: string(customerJourney.CustomerJourneyID()),
		TriggeredAt:       customerJourney.TriggeredAt(),
	})

	return nil
}
