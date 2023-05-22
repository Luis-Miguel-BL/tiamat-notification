package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/workflow/command"
)

type TriggerActionUsecase struct {
	customerRepo   repository.CustomerRepository
	segmentRepo    repository.SegmentRepository
	campaignRepo   repository.CampaignRepository
	triggerService service.TriggerService
	matcherService service.MatcherService
}

type NewTriggerActionUsecaseInput struct {
	CustomerRepo   repository.CustomerRepository
	SegmentRepo    repository.SegmentRepository
	CampaignRepo   repository.CampaignRepository
	TriggerService service.TriggerService
	MatcherService service.MatcherService
}

func NewTriggerActionUsecase(input NewTriggerActionUsecaseInput) *TriggerActionUsecase {
	return &TriggerActionUsecase{
		customerRepo:   input.CustomerRepo,
		segmentRepo:    input.SegmentRepo,
		campaignRepo:   input.CampaignRepo,
		triggerService: input.TriggerService,
		matcherService: input.MatcherService,
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
	campaignFilters, err := uc.getAllSegments(ctx, workspaceID, campaign.Filters())
	if err != nil {
		return err
	}

	err = uc.triggerService.TriggerAction(ctx, &customer, action, campaignFilters, campaign.CampaignID())
	if err != nil {
		return err
	}

	return nil
}

func (uc *TriggerActionUsecase) getAllSegments(ctx context.Context, workspaceID model.WorkspaceID, segmentIDs []model.SegmentID) (segments []model.Segment, err error) {
	for _, segmentID := range segmentIDs {
		segment, err := uc.segmentRepo.GetByID(ctx, segmentID, workspaceID)
		if err != nil {
			return segments, err
		}
		segments = append(segments, segment)
	}
	return segments, nil
}
