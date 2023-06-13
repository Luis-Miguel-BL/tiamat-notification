package journey

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
)

type StartJourneyService struct {
}

func NewStartJourneyService(repo repository.CustomerRepository) StartJourneyService {
	return StartJourneyService{}
}

func (s *StartJourneyService) StartJourney(ctx context.Context, customer *model.Customer, campaign model.Campaign, customerJourneys []model.Journey) (journey *model.Journey, err error) {
	startedJourney, alreadyStarted := alreadyStartedJourney(campaign.CampaignID(), customerJourneys)
	if alreadyStarted {
		if !campaign.MustBeTriggered(startedJourney.FinishedAt()) {
			return journey, nil
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

func alreadyStartedJourney(campaignID model.CampaignID, journeys []model.Journey) (journeyStarted model.Journey, alreadyStarted bool) {
	for _, journey := range journeys {
		if journey.CampaignID() == campaignID && journey.StartedAt().After(journey.StartedAt()) {
			return journey, true
		}
	}
	return journeyStarted, false
}
