package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service"
)

type MatchCustomerUsecase struct {
	customerRepo   repository.CustomerRepository
	segmentRepo    repository.SegmentRepository
	campaignRepo   repository.CampaignRepository
	matcherService service.MatcherService
}

func NewMatchCustomerUsecase(customerRepo repository.CustomerRepository, segmentRepo repository.SegmentRepository, campaignRepo repository.CampaignRepository, matcherService service.MatcherService) *MatchCustomerUsecase {
	return &MatchCustomerUsecase{
		customerRepo:   customerRepo,
		segmentRepo:    segmentRepo,
		campaignRepo:   campaignRepo,
		matcherService: matcherService,
	}
}

func (uc *MatchCustomerUsecase) MatchCustomer(ctx context.Context, command command.MatchCustomerCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	customerID := model.NewCustomerID(command.CustomerID)
	workspaceID := model.NewWorkspaceID(command.WorkspaceID)
	customer, err := uc.customerRepo.GetByID(ctx, customerID, workspaceID)
	if err != nil {
		return err
	}

	activeCampaigns, err := uc.campaignRepo.FindActiveCampaigns(ctx, workspaceID)
	if err != nil {
		return err
	}

	for _, activeCampaign := range activeCampaigns {
		isMatchWithTheTriggers, err := uc.matchAllSegments(ctx, workspaceID, activeCampaign.Triggers(), &customer)
		if err != nil {
			return err
		}
		if isMatchWithTheTriggers {

		}

	}

	return nil
}

func (uc *MatchCustomerUsecase) matchAllSegments(ctx context.Context, workspaceID model.WorkspaceID, segmentIDs []model.SegmentID, customer *model.Customer) (isMatchAll bool, err error) {
	isMatchAll = true
	for _, segmentID := range segmentIDs {
		segment, err := uc.segmentRepo.GetByID(ctx, segmentID, workspaceID)
		if err != nil {
			return isMatchAll, err
		}
		isMatchOne := uc.matcherService.MatchCustomerWithSegment(ctx, customer, segment)
		if !isMatchOne {
			isMatchAll = false
		}
	}
	return isMatchAll, nil
}