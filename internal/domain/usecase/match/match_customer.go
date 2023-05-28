package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/journey"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/matcher"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/match/command"
)

type MatchCustomerUsecase struct {
	customerRepo   repository.CustomerRepository
	segmentRepo    repository.SegmentRepository
	campaignRepo   repository.CampaignRepository
	matcherService matcher.CustomerMatcherService
	journeyService journey.StartJourneyService
}

type NewMatchCustomerUsecaseInput struct {
	CustomerRepo   repository.CustomerRepository
	SegmentRepo    repository.SegmentRepository
	CampaignRepo   repository.CampaignRepository
	MatcherService matcher.CustomerMatcherService
	JourneyService journey.StartJourneyService
}

func NewMatchCustomerUsecase(input NewMatchCustomerUsecaseInput) *MatchCustomerUsecase {
	return &MatchCustomerUsecase{
		customerRepo:   input.CustomerRepo,
		segmentRepo:    input.SegmentRepo,
		campaignRepo:   input.CampaignRepo,
		matcherService: input.MatcherService,
		journeyService: input.JourneyService,
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

	campaigns, err := uc.campaignRepo.FindAll(ctx, workspaceID)
	if err != nil {
		return err
	}

	for _, campaign := range campaigns {
		if !campaign.IsActive() {
			continue
		}
		isMatchWithTheTriggers, err := uc.matchAllSegments(ctx, workspaceID, campaign.Triggers(), &customer)
		if err != nil {
			return err
		}
		if isMatchWithTheTriggers {
			err = uc.journeyService.StartJourney(ctx, &customer, campaign)
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
