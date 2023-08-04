package consumer

import (
	"context"

	usecase "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging/marshaller"
)

type CustomerEventOccurredConsumer struct {
	uc  usecase.MatchCustomerUsecase
	log logger.Logger
}

func NewCustomerEventOccurredConsumer(matchCustomerUsecase usecase.MatchCustomerUsecase, log logger.Logger) *CustomerEventOccurredConsumer {
	return &CustomerEventOccurredConsumer{
		uc:  matchCustomerUsecase,
		log: log,
	}
}

func (c *CustomerEventOccurredConsumer) Consume(ctx context.Context, eventStr string) (err error) {
	event, err := marshaller.UnmarshalDomainEvent[*event.CustomerEventOccurredEvent](eventStr)
	if err != nil {
		return err
	}

	err = c.uc.MatchCustomer(ctx, input.MatchCustomerInput{
		CustomerID:  event.CustomerID,
		WorkspaceID: event.WorkspaceID,
	})
	if err != nil {
		return err
	}

	return nil
}
