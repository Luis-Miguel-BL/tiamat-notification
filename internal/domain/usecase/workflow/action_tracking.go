package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/workflow/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type ActionTrackingUsecase struct {
	customerRepo repository.CustomerRepository
	campaignRepo repository.CampaignRepository
}

type NewActionTrackingUsecaseInput struct {
	CustomerRepo repository.CustomerRepository
	CampaignRepo repository.CampaignRepository
}

func NewActionTrackingUsecase(input NewActionTrackingUsecaseInput) *ActionTrackingUsecase {
	return &ActionTrackingUsecase{
		customerRepo: input.CustomerRepo,
		campaignRepo: input.CampaignRepo,
	}
}

func (uc *ActionTrackingUsecase) TriggeredActionTracking(ctx context.Context, command command.TriggeredActionTrackingCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	stepJourneyID := model.StepJourneyID(command.StepJourneyID)
	eventSlug, err := vo.NewSlug(command.EventSlug)
	if err != nil {
		return err
	}
	trackingData, err := vo.NewCustomAttributes(command.TrackingData)
	if err != nil {
		return err
	}
	stepJourney, err := uc.customerRepo.GetStepJourney(ctx, workspaceID, stepJourneyID)
	if err != nil {
		return err
	}
	stepJourney.AppendTrackingEvent(eventSlug, trackingData)

	err = uc.customerRepo.SaveStepJourney(ctx, stepJourney)
	if err != nil {
		return err
	}

	return nil
}
