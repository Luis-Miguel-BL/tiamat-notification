package service

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type MatcherService interface {
	MatchCustomerWithSegments(ctx context.Context, customer *model.Customer, segments []model.Segment) (isMatch bool)
}

type customerSegmentation struct {
	repo repository.CustomerRepository
}

func NewMatcherService(repo repository.CustomerRepository) MatcherService {
	return &customerSegmentation{
		repo: repo,
	}
}

func (s *customerSegmentation) MatchCustomerWithSegments(ctx context.Context, customer *model.Customer, segments []model.Segment) (isMatch bool) {
	isMatch = true
	for _, segment := range segments {
	conditionLoop:
		for _, condition := range segment.Conditions {
			if !condition.IsMatch(customer.Serialize()) {
				isMatch = false
				break conditionLoop
			}
		}

		currentSegment := model.NewCurrentSegment(
			model.NewCurrentSegmentInput{
				CurrentSegmentID: model.NewCurrentSegmentID(util.NewUUID()),
				SegmentID:        segment.SegmentID,
			},
		)
		customer.AppendCurrentSegment(*currentSegment)
	}

	return isMatch
}
