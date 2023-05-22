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
	customerJourneyID := model.CustomerJourneyID(command.CustomerJourneyID)
	eventSlug, err := vo.NewSlug(command.EventSlug)
	if err != nil {
		return err
	}
	trackingData, err := vo.NewCustomAttributes(command.TrackingData)
	if err != nil {
		return err
	}
	customerJourney, err := uc.customerRepo.GetCustomerJourney(ctx, workspaceID, customerJourneyID)
	if err != nil {
		return err
	}
	customerJourney.AppendTrackingEvent(eventSlug, trackingData)

	err = uc.customerRepo.SaveCustomerJourney(ctx, customerJourney)
	if err != nil {
		return err
	}

	return nil
}
