package main

import (
	"context"

	consumer "github.com/Luis-Miguel-BL/tiamat-notification/internal/api/consumer/match"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/event"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging/event_bus"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/persistence/repository"
)

func main() {
	ctx := context.Background()
	config := config.LoadConfig()
	log := logger.NewZerologger(config.AppName)
	log.SetLevel(logger.Debug)
	eventBus := event_bus.NewRedisClient(config.Redis, log, config.AppName)
	eventDispatcher := messaging.NewAggregateEventDispatcher(eventBus)
	repoManager, err := repository.NewRepositoryManager(ctx, *eventDispatcher, config.DBConfig, log)
	if err != nil {
		panic(err)
	}
	gatewayManager := gateway.GatewayManager{}
	usecaseManager := usecase.NewUsecaseManager(repoManager, gatewayManager)
	consumer := consumer.NewCustomerChangedEventConsumer(usecaseManager.MatchCustomer, log)
	log.Info("start consumer ...")
	events := make(chan domain.DomainEvent)
	err = eventBus.Subscribe(ctx, events, event.CustomerCreatedEventType, event.CustomerUpdatedEventType, event.CustomerEventOccurredEventType)
	if err != nil {
		panic(err)
	}

	for event := range events {
		log.Debugf("send event to consumer %+v", event)
		err = consumer.Consume(ctx, event)
		if err != nil {
			log.Errorf("error at consume event: %s", err.Error())
		}
	}

}
