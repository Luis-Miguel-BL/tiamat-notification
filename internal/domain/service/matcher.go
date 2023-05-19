package service

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
)

type MatcherService interface {
	MatchCustomerWithSegment(ctx context.Context, customer *model.Customer, segment model.Segment) (isMatch bool)
}

type matcherService struct {
}

func NewMatcherService(repo repository.CustomerRepository) MatcherService {
	return &matcherService{}
}

func (s *matcherService) MatchCustomerWithSegment(ctx context.Context, customer *model.Customer, segment model.Segment) (isMatch bool) {
	for _, condition := range segment.Conditions() {
		if !condition.IsMatch(customer.Serialize()) {
			return false
		}
	}

	satisfiedSegment, err := model.NewSatisfiedSegment(
		model.NewSatisfiedSegmentInput{
			CustomerID:  customer.CustomerID(),
			WorkspaceID: customer.WorkspaceID(),
			SegmentID:   segment.SegmentID,
		},
	)
	if err != nil {
		return false
	}
	customer.AppendSatisfiedSegment(*satisfiedSegment)

	return true
}
