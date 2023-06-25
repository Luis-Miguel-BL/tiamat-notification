package consumers

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	usecase "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/match/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging/marshaller"
)

type ChangedCustomerConsumer struct {
	matchCustomerUsecase usecase.MatchCustomerUsecase
	log                  logger.Logger
}

func NewChangedCustomerConsumer(matchCustomerUsecase usecase.MatchCustomerUsecase, log logger.Logger) *ChangedCustomerConsumer {
	return &ChangedCustomerConsumer{
		matchCustomerUsecase: matchCustomerUsecase,
		log:                  log,
	}
}

func (c *ChangedCustomerConsumer) Consume(ctx context.Context, eventType string, eventStr string) (err error) {
	var customerID string
	var workspaceID string
	switch domain.EventType(eventType) {
	case event.CustomerCreatedEventType:
		domainEvent, err := marshaller.UnmarshalDomainEvent[event.CustomerCreatedEvent](eventStr)
		if err != nil {
			return err
		}
		customerID = domainEvent.CustomerID
		workspaceID = domainEvent.WorkspaceID
	case event.CustomerEventOccurredEventType:
		domainEvent, err := marshaller.UnmarshalDomainEvent[event.CustomerEventOccurredEvent](eventStr)
		if err != nil {
			return err
		}
		customerID = domainEvent.CustomerID
		workspaceID = domainEvent.WorkspaceID
	default:
		return errors.NewInvalidParamError("event_type")
	}

	err = c.matchCustomerUsecase.MatchCustomer(ctx, input.MatchCustomerInput{
		CustomerID:  customerID,
		WorkspaceID: workspaceID,
	})
	if err != nil {
		return err
	}

	return nil
}
