package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/journey"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/workflow/command"
)

type TriggerActionUsecase struct {
	customerRepo   repository.CustomerRepository
	segmentRepo    repository.SegmentRepository
	campaignRepo   repository.CampaignRepository
	triggerService journey.TriggerStepJourneyService
}

type NewTriggerActionUsecaseInput struct {
	CustomerRepo   repository.CustomerRepository
	SegmentRepo    repository.SegmentRepository
	CampaignRepo   repository.CampaignRepository
	TriggerService journey.TriggerStepJourneyService
}

func NewTriggerActionUsecase(input NewTriggerActionUsecaseInput) *TriggerActionUsecase {
	return &TriggerActionUsecase{
		customerRepo:   input.CustomerRepo,
		segmentRepo:    input.SegmentRepo,
		campaignRepo:   input.CampaignRepo,
		triggerService: input.TriggerService,
	}
}

func (uc *TriggerActionUsecase) TriggerAction(ctx context.Context, command command.TriggerActionCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	customerID := model.CustomerID(command.CustomerID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	campaignID := model.CampaignID(command.CampaignID)
	actionID := model.ActionID(command.ActionID)
	customer, err := uc.customerRepo.GetByID(ctx, customerID, workspaceID)
	if err != nil {
		return err
	}
	defer func() {
		uc.customerRepo.Save(ctx, customer)
	}()

	campaign, err := uc.campaignRepo.GetByID(ctx, campaignID, workspaceID)
	if err != nil {
		return err
	}
	action, err := uc.campaignRepo.GetActionByID(ctx, campaignID, actionID, workspaceID)
	if err != nil {
		return err
	}

	err = uc.triggerService.TriggerStepJourney(ctx, &customer, action, campaign)
	if err != nil {
		return err
	}

	return nil
}
