package consumer

import (
	"context"

	usecase "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging/marshaller"
)

type CustomerCreatedConsumer struct {
	uc  usecase.MatchCustomerUsecase
	log logger.Logger
}

func NewCustomerCreatedConsumer(matchCustomerUsecase usecase.MatchCustomerUsecase, log logger.Logger) *CustomerCreatedConsumer {
	return &CustomerCreatedConsumer{
		uc:  matchCustomerUsecase,
		log: log,
	}
}

func (c *CustomerCreatedConsumer) Consume(ctx context.Context, eventStr string) (err error) {
	event, err := marshaller.UnmarshalDomainEvent[*event.CustomerCreatedEvent](eventStr)
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
