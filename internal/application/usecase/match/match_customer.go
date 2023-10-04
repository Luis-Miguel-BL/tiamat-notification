package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/journey"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/matcher"
)

type MatchCustomerUsecase struct {
	customerRepo   repository.CustomerRepository
	segmentRepo    repository.SegmentRepository
	campaignRepo   repository.CampaignRepository
	journeyRepo    repository.JourneyRepository
	matcherService matcher.CustomerMatcherService
	journeyService journey.StartJourneyService
}

func NewMatchCustomerUsecase(customerRepo repository.CustomerRepository, segmentRepo repository.SegmentRepository, campaignRepo repository.CampaignRepository, journeyRepo repository.JourneyRepository) *MatchCustomerUsecase {
	return &MatchCustomerUsecase{
		customerRepo:   customerRepo,
		segmentRepo:    segmentRepo,
		campaignRepo:   campaignRepo,
		journeyRepo:    journeyRepo,
		matcherService: matcher.NewCustomerMatcherService(),
		journeyService: journey.NewStartJourneyService(),
	}
}

func (uc *MatchCustomerUsecase) MatchCustomer(ctx context.Context, input input.MatchCustomerInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	customerID := model.CustomerID(input.CustomerID)
	workspaceID := model.WorkspaceID(input.WorkspaceID)
	customer, err := uc.customerRepo.GetByID(ctx, customerID, workspaceID)
	if err != nil {
		return err
	}
	defer func() {
		uc.customerRepo.Save(ctx, *customer)
	}()

	customerJourneys, err := uc.journeyRepo.FindByCustomerID(ctx, customer.WorkspaceID(), customer.CustomerID())
	if err != nil {
		return err
	}

	campaigns, err := uc.campaignRepo.FindAll(ctx, workspaceID)
	if err != nil {
		return err
	}

	for _, campaign := range campaigns {
		if !campaign.IsActive() {
			continue
		}
		isMatchWithTheTriggers, err := uc.matchAllSegments(ctx, workspaceID, campaign.Triggers(), customer)
		if err != nil {
			return err
		}
		if isMatchWithTheTriggers {
			newJourney, err := uc.journeyService.StartJourney(ctx, customer, campaign, customerJourneys)
			if err != nil {
				return err
			}

			uc.journeyRepo.Save(ctx, *newJourney)
			if err != nil {
				return err
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
