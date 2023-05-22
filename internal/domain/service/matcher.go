package service

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
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

	satisfiedSegment, err := model.NewCustomerSegment(
		model.NewCustomerSegmentInput{
			CustomerID:  customer.CustomerID(),
			WorkspaceID: customer.WorkspaceID(),
			SegmentID:   segment.SegmentID(),
		},
	)
	if err != nil {
		return false
	}
	customer.AppendCustomerSegment(*satisfiedSegment)

	customer.AggregateRoot.AppendEvent(event.CustomerMatched{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerEventOccurredEventType,
			OccurredAt:    satisfiedSegment.MatchedAt(),
			AggregateType: customer.AggregateType(),
			AggregateID:   customer.AggregateID(),
		}),
		CustomerID:  string(customer.CustomerID()),
		WorkspaceID: string(customer.WorkspaceID()),
		SegmentID:   string(satisfiedSegment.SegmentID()),
		MatchedAt:   satisfiedSegment.MatchedAt(),
	})

	return true
}
