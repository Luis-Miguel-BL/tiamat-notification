package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
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

		isMatchWithTheTriggers := uc.matcherService.MatchCustomerWithSegments(ctx, &customer)
	}

	eventSlug, err := vo.NewSlug(command.Slug)
	if err != nil {
		return err
	}
	customAttr, err := vo.NewCustomAttributes(command.CustomAttributes)
	if err != nil {
		return err
	}

	customerEventToCreate, err := model.NewCustomerEvent(
		model.NewCustomerEventInput{
			CustomerEventID:  customerEventID,
			Slug:             eventSlug,
			CustomAttributes: customAttr,
		},
	)
	if err != nil {
		return err
	}

	err = uc.repo.MatchCustomer(ctx, customerID, workspaceID, *customerEventToCreate)

	return err
}

func matchAllSegments(ctx context.Context)
