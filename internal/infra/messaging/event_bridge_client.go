package messaging

import (
	"context"
	"encoding/json"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
)

type EventBridgeClient struct {
	cfg    config.EventBridgeConfig
	log    logger.Logger
	conn   *eventbridge.EventBridge
	source string
}

func NewEventBridgeClient(cfg config.EventBridgeConfig, log logger.Logger, source string) *EventBridgeClient {
	session := session.Must(session.NewSession())
	conn := eventbridge.New(session, aws.NewConfig().WithRegion(cfg.Region))

	return &EventBridgeClient{
		cfg:    cfg,
		log:    log,
		conn:   conn,
		source: source,
	}
}

func (e *EventBridgeClient) Publish(ctx context.Context, event domain.DomainEvent) (err error) {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = e.conn.PutEvents(&eventbridge.PutEventsInput{
		Entries: []*eventbridge.PutEventsRequestEntry{{
			EventBusName: aws.String(e.cfg.EventBusName),
			Source:       aws.String(e.source),
			DetailType:   aws.String(string(event.EventType())),
			Detail:       aws.String(string(eventJson)),
		}},
	})
	if err != nil {
		return err
	}
	return nil
}
