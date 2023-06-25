package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/workflow/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/journey"
)

type TriggerActionUsecase struct {
	customerRepo   repository.CustomerRepository
	segmentRepo    repository.SegmentRepository
	campaignRepo   repository.CampaignRepository
	journeyRepo    repository.JourneyRepository
	triggerService journey.TriggerStepJourneyService
}

func NewTriggerActionUsecase(customerRepo repository.CustomerRepository, segmentRepo repository.SegmentRepository, campaignRepo repository.CampaignRepository, journeyRepo repository.JourneyRepository, gatewayManager gateway.GatewayManager) *TriggerActionUsecase {
	return &TriggerActionUsecase{
		customerRepo:   customerRepo,
		segmentRepo:    segmentRepo,
		campaignRepo:   campaignRepo,
		journeyRepo:    journeyRepo,
		triggerService: journey.NewTriggerStepJourneyService(gatewayManager),
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

	err = uc.triggerService.TriggerStepJourney(ctx, &journey, *customer, action, campaign)
	if err != nil {
		return err
	}

	return nil
}
