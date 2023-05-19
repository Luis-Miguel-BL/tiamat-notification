package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/match/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service"
)

type MatchCustomerUsecase struct {
	customerRepo   repository.CustomerRepository
	segmentRepo    repository.SegmentRepository
	campaignRepo   repository.CampaignRepository
	matcherService service.MatcherService
	triggerService service.TriggerService
}

type NewMatchCustomerUsecaseInput struct {
	CustomerRepo   repository.CustomerRepository
	SegmentRepo    repository.SegmentRepository
	CampaignRepo   repository.CampaignRepository
	MatcherService service.MatcherService
	TriggerService service.TriggerService
}

func NewMatchCustomerUsecase(input NewMatchCustomerUsecaseInput) *MatchCustomerUsecase {
	return &MatchCustomerUsecase{
		customerRepo:   input.CustomerRepo,
		segmentRepo:    input.SegmentRepo,
		campaignRepo:   input.CampaignRepo,
		matcherService: input.MatcherService,
		triggerService: input.TriggerService,
	}
}

func (uc *MatchCustomerUsecase) MatchCustomer(ctx context.Context, command command.MatchCustomerCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	customerID := model.CustomerID(command.CustomerID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)
	customer, err := uc.customerRepo.GetByID(ctx, customerID, workspaceID)
	if err != nil {
		return err
	}
	defer func() {
		uc.customerRepo.Save(ctx, customer)
	}()

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
			err = uc.triggerService.TriggerCampaign(ctx, &customer, activeCampaign)
			if err != nil {
				continue
			}
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
