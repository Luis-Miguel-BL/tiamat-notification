package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/workflow/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type ActionTrackingUsecase struct {
	journeyRepo  repository.JourneyRepository
	campaignRepo repository.CampaignRepository
}

func NewActionTrackingUsecase(journeyRepo repository.JourneyRepository, campaignRepo repository.CampaignRepository) *ActionTrackingUsecase {
	return &ActionTrackingUsecase{
		journeyRepo:  journeyRepo,
		campaignRepo: campaignRepo,
	}
}

func (uc *ActionTrackingUsecase) TriggeredActionTracking(ctx context.Context, input input.TriggeredActionTrackingInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(input.WorkspaceID)
	stepJourneyID := model.StepJourneyID(input.StepJourneyID)
	eventSlug, err := vo.NewSlug(input.EventSlug)
	if err != nil {
		return err
	}
	trackingData, err := vo.NewCustomAttributes(input.TrackingData)
	if err != nil {
		return err
	}
	stepJourney, err := uc.journeyRepo.GetStepJourney(ctx, workspaceID, stepJourneyID)
	if err != nil {
		return err
	}
	stepJourney.AppendTrackingEvent(eventSlug, trackingData)

	err = uc.journeyRepo.SaveStepJourney(ctx, stepJourney)
	if err != nil {
		return err
	}

	return nil
}
