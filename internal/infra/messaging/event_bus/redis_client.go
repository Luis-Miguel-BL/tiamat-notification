package event_bus

import (
	"context"
	"encoding/json"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging/marshaller"
	go_redis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	cfg    config.RedisConfig
	log    logger.Logger
	client *go_redis.Client
	source string
}

func NewRedisClient(cfg config.RedisConfig, log logger.Logger, source string) *RedisClient {
	client := go_redis.NewClient(&go_redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return &RedisClient{
		cfg:    cfg,
		log:    log,
		client: client,
		source: source,
	}
}

func (e *RedisClient) Publish(ctx context.Context, event domain.DomainEvent) (err error) {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = e.client.Publish(ctx, string(event.EventType()), string(eventJson)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (e *RedisClient) Subscribe(ctx context.Context, messages chan domain.DomainEvent, eventTypes ...domain.EventType) (err error) {
	var channels []string
	for _, eventType := range eventTypes {
		channels = append(channels, string(eventType))
	}
	redisSubscription := e.client.Subscribe(ctx, channels...)

	go func() {
		for message := range redisSubscription.Channel() {
			e.log.Debugf("receive message: %s", message.String())
			domainEvent, err := marshaller.UnmarshalDomainEvent(domain.EventType(message.Channel), message.Payload)
			if err != nil {
				e.log.Warnf("error at unmarshal message: %s", err.Error())
			}
			e.log.Debugf("parsed event: %s", domainEvent)
			messages <- domainEvent
		}
	}()

	return err
}
