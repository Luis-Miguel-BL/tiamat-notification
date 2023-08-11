package handler

import (
	"context"
	"fmt"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api"
	match_consumer "github.com/Luis-Miguel-BL/tiamat-notification/internal/api/consumer/match"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging/marshaller"
	"github.com/aws/aws-lambda-go/events"
)

func (h *ServerlessHandler) sqsHandle(ctx context.Context, event events.SQSEvent) (response events.SQSEventResponse, err error) {
	var consumer api.EventConsumer
	switch h.cfg.AppName {
	case "customer-change-event-consumer":
		consumer = match_consumer.NewCustomerChangedEventConsumer(h.usecaseManager.MatchCustomer, h.log)
	default:
		return response, fmt.Errorf("app-name not found: %s", h.cfg.AppName)
	}

	for _, record := range event.Records {
		eventType, err := marshaller.UnmarshalDomainEventType(record.Body)
		if err != nil {
			response.BatchItemFailures = append(response.BatchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: record.MessageId})
			continue
		}

		err = consumer.Consume(ctx, eventType, record.Body)
		if err != nil {
			response.BatchItemFailures = append(response.BatchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: record.MessageId})
			continue
		}
	}

	return response, nil
}
