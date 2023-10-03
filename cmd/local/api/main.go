package main

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging/event_bus/event_bridge"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/persistence/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/server"
)

func main() {
	ctx := context.Background()
	config := config.LoadConfig()
	log := logger.NewZerologger(config.AppName)
	log.SetLevel(logger.Debug)
	eventBus := event_bridge.NewEventBridgeClient(config.EventBridge, log, config.AppName)
	eventDispatcher := messaging.NewAggregateEventDispatcher(eventBus)
	repoManager, err := repository.NewRepositoryManager(ctx, *eventDispatcher, config.DBConfig, log)
	if err != nil {
		panic(err)
	}
	gatewayManager := gateway.GatewayManager{}
	usecaseManager := usecase.NewUsecaseManager(repoManager, gatewayManager)

	srv, err := server.NewServer(log, config, *usecaseManager)
	if err != nil {
		panic(err)
	}

	srv.Run(ctx)
}
