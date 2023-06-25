package main

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/cmd/serverless/handler"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/persistence/repository"
	rgt "github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/registry"
	"github.com/aws/aws-lambda-go/events"
)

var ctx context.Context
var cfg *config.Config
var log logger.Logger
var usecaseManager usecase.UsecaseManager

func init() {
	ctx = context.Background()
	cfg = config.LoadConfig()
	log = logger.NewZerologger(cfg.AppName)
	eventBus := messaging.NewEventBridgeClient(cfg.EventBridge, log, cfg.AppName)
	dispatcher := messaging.NewAggregateEventDispatcher(eventBus)
	repositoryManager, err := repository.NewRepositoryManager(ctx, *dispatcher, cfg.DBConfig, log)
	if err != nil {
		panic(err)
	}

	usecaseManager = *usecase.NewUsecaseManager(repositoryManager, gateway.GatewayManager{})
}

func main() {
	registry := rgt.NewRegistry()

	registry.Provide("usecases", usecaseManager)
	registry.Provide("logger", log)
	registry.Provide("config", cfg)

	ctx = context.WithValue(ctx, "registry", registry)

	switch cfg.EntryPoint {
	case "event-bridge":
		// lambda.StartWithOptions(handler.EventBridgeHandle, lambda.WithContext(ctx))
		handler.EventBridgeHandle(ctx, events.CloudWatchEvent{})
	default:
		panic("entry point not found")
	}

}
