package service

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type MatcherService interface {
	MatchCustomerWithSegment(ctx context.Context, customer *model.Customer, segment model.Segment) (isMatch bool)
}

type customerSegmentation struct {
	repo repository.CustomerRepository
}

func NewMatcherService(repo repository.CustomerRepository) MatcherService {
	return &customerSegmentation{
		repo: repo,
	}
}

func (s *customerSegmentation) MatchCustomerWithSegment(ctx context.Context, customer *model.Customer, segment model.Segment) (isMatch bool) {
	for _, condition := range segment.Conditions {
		if !condition.IsMatch(customer.Serialize()) {
			return false
		}
	}

	currentSegment := model.NewCurrentSegment(
		model.NewCurrentSegmentInput{
			CurrentSegmentID: model.NewCurrentSegmentID(util.NewUUID()),
			SegmentID:        segment.SegmentID,
		},
	)
	customer.AppendCurrentSegment(*currentSegment)

	return true
}
