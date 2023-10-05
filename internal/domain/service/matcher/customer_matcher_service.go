package matcher

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/service/matcher/match_condition"
)

type CustomerMatcherService struct {
}

func NewCustomerMatcherService() CustomerMatcherService {
	return CustomerMatcherService{}
}

func (s *CustomerMatcherService) MatchCustomerWithSegment(ctx context.Context, customer *model.Customer, segment model.Segment) (isMatch bool) {
	serializedCustomer := customer.Serialize()
	for _, condition := range segment.Conditions() {
		if !match_condition.Match(condition, serializedCustomer) {
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

	customer.AggregateRoot.AppendEvent(event.CustomerMatchedEvent{
		DomainEventBase: domain.NewDomainEventBase(domain.NewDomainEventBaseInput{
			EventType:     event.CustomerMatchedEventType,
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
