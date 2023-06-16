package journey

import (
	"context"
	"fmt"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
)

type StartJourneyService struct {
}

func NewStartJourneyService(repo repository.CustomerRepository) StartJourneyService {
	return StartJourneyService{}
}

func (s *StartJourneyService) StartJourney(ctx context.Context, customer *model.Customer, campaign model.Campaign, customerJourneys []model.Journey) (journey *model.Journey, err error) {
	previousJourney, find := findPreviousJourney(campaign.CampaignID(), customerJourneys)
	if find {
		if !previousJourney.IsFinished() {
			return journey, domain.DomainError(fmt.Errorf("journey already started"))
		}
		if !campaign.MustBeRetriggered(previousJourney.FinishedAt()) {
			return journey, domain.DomainError(fmt.Errorf("journey cannot be started yet"))
		}
	}

	firstAction, err := campaign.Action(campaign.FirstActionID())
	if err != nil {
		return journey, err
	}
	journey, err = model.NewJourney(model.NewJourneyInput{
		WorkspaceID: customer.WorkspaceID(),
		CustomerID:  customer.CustomerID(),
		CampaignID:  campaign.CampaignID(),
	})
	if err != nil {
		return journey, err
	}

	err = journey.AppendNextStepJourney(firstAction.ActionID())
	if err != nil {
		return journey, err
	}

	return journey, nil
}

func findPreviousJourney(campaignID model.CampaignID, journeys []model.Journey) (previousJourney model.Journey, find bool) {
	for _, journey := range journeys {
		if journey.CampaignID() == campaignID && journey.StartedAt().After(journey.StartedAt()) {
			return journey, true
		}
	}
	return previousJourney, false
}
