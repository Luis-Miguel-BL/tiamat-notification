package consumer

import (
	"context"
	"fmt"

	usecase "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	domain_event "github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
)

type CustomerChangedEventConsumer struct {
	uc  *usecase.MatchCustomerUsecase
	log logger.Logger
}

func NewCustomerChangedEventConsumer(matchCustomerUsecase *usecase.MatchCustomerUsecase, log logger.Logger) *CustomerChangedEventConsumer {
	return &CustomerChangedEventConsumer{
		uc:  matchCustomerUsecase,
		log: log,
	}
}

func (c *CustomerChangedEventConsumer) Consume(ctx context.Context, eventOccurred domain.DomainEvent) (err error) {
	var ucInput input.MatchCustomerInput

	switch event := eventOccurred.(type) {
	case domain_event.CustomerCreatedEvent:
		ucInput.CustomerID = event.CustomerID
		ucInput.WorkspaceID = event.WorkspaceID
	case domain_event.CustomerUpdatedEvent:
		ucInput.CustomerID = event.CustomerID
		ucInput.WorkspaceID = event.WorkspaceID
	case domain_event.CustomerEventOccurredEvent:
		ucInput.CustomerID = event.CustomerID
		ucInput.WorkspaceID = event.WorkspaceID
	default:
		return fmt.Errorf("invalid event-type %s", eventOccurred.EventType())
	}

	err = c.uc.MatchCustomer(ctx, ucInput)
	if err != nil {
		return err
	}

	return nil
}
