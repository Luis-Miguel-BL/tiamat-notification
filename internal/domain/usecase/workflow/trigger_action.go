package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/journey"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/workflow/input"
)

type TriggerActionUsecase struct {
	customerRepo   repository.CustomerRepository
	segmentRepo    repository.SegmentRepository
	campaignRepo   repository.CampaignRepository
	journeyRepo    repository.JourneyRepository
	triggerService journey.TriggerStepJourneyService
}

type NewTriggerActionUsecaseInput struct {
	CustomerRepo   repository.CustomerRepository
	SegmentRepo    repository.SegmentRepository
	CampaignRepo   repository.CampaignRepository
	JourneyRepo    repository.JourneyRepository
	TriggerService journey.TriggerStepJourneyService
}

func NewTriggerActionUsecase(input NewTriggerActionUsecaseInput) *TriggerActionUsecase {
	return &TriggerActionUsecase{
		customerRepo:   input.CustomerRepo,
		segmentRepo:    input.SegmentRepo,
		campaignRepo:   input.CampaignRepo,
		journeyRepo:    input.JourneyRepo,
		triggerService: input.TriggerService,
	}
}

func (uc *TriggerActionUsecase) TriggerAction(ctx context.Context, input input.TriggerActionInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(input.WorkspaceID)
	journeyID := model.JourneyID(input.JourneyID)
	customerID := model.CustomerID(input.CustomerID)
	campaignID := model.CampaignID(input.CampaignID)
	actionID := model.ActionID(input.ActionID)
	journey, err := uc.journeyRepo.GetByID(ctx, journeyID, workspaceID)
	if err != nil {
		return err
	}
	defer func() {
		uc.journeyRepo.Save(ctx, journey)
	}()

	customer, err := uc.customerRepo.GetByID(ctx, customerID, workspaceID)
	if err != nil {
		return err
	}

	campaign, err := uc.campaignRepo.GetByID(ctx, campaignID, workspaceID)
	if err != nil {
		return err
	}
	action, err := uc.campaignRepo.GetActionByID(ctx, campaignID, actionID, workspaceID)
	if err != nil {
		return err
	}

	err = uc.triggerService.TriggerStepJourney(ctx, &journey, customer, action, campaign)
	if err != nil {
		return err
	}

	return nil
}
