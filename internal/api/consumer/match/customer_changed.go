package consumer

import (
	"context"
	"fmt"

	usecase "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging/marshaller"
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

func (c *CustomerChangedEventConsumer) Consume(ctx context.Context, eventType domain.EventType, eventStr string) (err error) {
	var ucInput input.MatchCustomerInput

	switch eventType {
	case event.CustomerCreatedEventType:
		event, err := marshaller.UnmarshalDomainEvent[*event.CustomerCreatedEvent](eventStr)
		if err != nil {
			return err
		}
		ucInput.CustomerID = event.CustomerID
		ucInput.WorkspaceID = event.WorkspaceID
	case event.CustomerUpdatedEventType:
		event, err := marshaller.UnmarshalDomainEvent[*event.CustomerUpdatedEvent](eventStr)
		if err != nil {
			return err
		}
		ucInput.CustomerID = event.CustomerID
		ucInput.WorkspaceID = event.WorkspaceID
	case event.CustomerEventOccurredEventType:
		event, err := marshaller.UnmarshalDomainEvent[*event.CustomerEventOccurredEvent](eventStr)
		if err != nil {
			return err
		}
		ucInput.CustomerID = event.CustomerID
		ucInput.WorkspaceID = event.WorkspaceID
	default:
		return fmt.Errorf("invalid event-type %s", eventType)
	}

	err = c.uc.MatchCustomer(ctx, ucInput)
	if err != nil {
		return err
	}

	return nil
}
