package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/trigger/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type TriggeredActionTrackingUsecase struct {
	customerRepo repository.CustomerRepository
	campaignRepo repository.CampaignRepository
}

type NewTriggeredActionTrackingUsecaseInput struct {
	CustomerRepo repository.CustomerRepository
	CampaignRepo repository.CampaignRepository
}

func NewTriggeredActionTrackingUsecase(input NewTriggeredActionTrackingUsecaseInput) *TriggeredActionTrackingUsecase {
	return &TriggeredActionTrackingUsecase{
		customerRepo: input.CustomerRepo,
		campaignRepo: input.CampaignRepo,
	}
}

func (uc *TriggeredActionTrackingUsecase) TriggeredActionTracking(ctx context.Context, command command.TriggeredActionTrackingCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.NewWorkspaceID(command.WorkspaceID)
	actionTriggeredID := model.NewActionTriggeredID(command.ActionTriggeredID)
	eventSlug, err := vo.NewSlug(command.EventSlug)
	if err != nil {
		return err
	}
	trackingData, err := vo.NewCustomAttributes(command.TrackingData)
	if err != nil {
		return err
	}
	actionTriggered, err := uc.customerRepo.GetActionTriggered(ctx, workspaceID, actionTriggeredID)
	if err != nil {
		return err
	}
	actionTriggered.AppendTrackingEvent(eventSlug, trackingData)

	err = uc.customerRepo.SaveActionTriggered(ctx, actionTriggered)
	if err != nil {
		return err
	}

	return nil
}
