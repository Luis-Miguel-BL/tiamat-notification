package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service"
)

type TriggerActionUsecase struct {
	customerRepo         repository.CustomerRepository
	segmentRepo          repository.SegmentRepository
	campaignRepo         repository.CampaignRepository
	actionHandlerService service.ActionHandlerService
	matcherService       service.MatcherService
}

type NewTriggerActionUsecaseInput struct {
	CustomerRepo         repository.CustomerRepository
	SegmentRepo          repository.SegmentRepository
	CampaignRepo         repository.CampaignRepository
	ActionHandlerService service.ActionHandlerService
	MatcherService       service.MatcherService
}

func NewTriggerActionUsecase(input NewTriggerActionUsecaseInput) *TriggerActionUsecase {
	return &TriggerActionUsecase{
		customerRepo:         input.CustomerRepo,
		segmentRepo:          input.SegmentRepo,
		campaignRepo:         input.CampaignRepo,
		actionHandlerService: input.ActionHandlerService,
		matcherService:       input.MatcherService,
	}
}

func (uc *TriggerActionUsecase) TriggerAction(ctx context.Context, command command.TriggerActionCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	customerID := model.NewCustomerID(command.CustomerID)
	workspaceID := model.NewWorkspaceID(command.WorkspaceID)
	campaignID := model.NewCampaignID(command.CampaignID)
	actionID := model.NewActionID(command.ActionID)
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

	err = uc.actionHandlerService.HandleAction(ctx, &customer, action, campaignFilters)
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
