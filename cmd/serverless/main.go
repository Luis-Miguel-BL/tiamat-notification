package main

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/cmd/serverless/handler"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/gateway"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/persistence/repository"
	"github.com/aws/aws-lambda-go/lambda"
)

var ctx context.Context
var cfg *config.Config
var log logger.Logger
var usecaseManager *usecase.UsecaseManager

func main() {
	ctx = context.Background()

	cfg = config.LoadConfig()
	log = logger.NewZerologger(cfg.AppName)
	eventBus := messaging.NewEventBridgeClient(cfg.EventBridge, log, cfg.AppName)
	dispatcher := messaging.NewAggregateEventDispatcher(eventBus)
	repositoryManager, err := repository.NewRepositoryManager(ctx, *dispatcher, cfg.DBConfig, log)
	if err != nil {
		panic(err)
	}
	usecaseManager = usecase.NewUsecaseManager(repositoryManager, gateway.GatewayManager{})

	handlerManager := handler.NewServerlessHandler(cfg, log, usecaseManager)

	lambda.StartWithOptions(handlerManager, lambda.WithContext(ctx))

}
