package handler

import (
	"context"
	"fmt"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api/consumers"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/registry"
	"github.com/aws/aws-lambda-go/events"
)

func EventBridgeHandle(ctx context.Context, event events.CloudWatchEvent) error {
	registry := ctx.Value("registry").(*registry.Registry)

	log := registry.Inject("logger").(*logger.ZeroLogger)
	cfg := registry.Inject("config").(*config.Config)
	usecases := registry.Inject("usecases").(*usecase.UsecaseManager)

	var consumer consumers.EventConsumer
	switch cfg.AppName {
	case "customer-change-event-consumer":
		consumer = consumers.NewChangedCustomerConsumer(*usecases.MatchCustomer, *log)
	default:
		return fmt.Errorf("app-name not found: %s", cfg.AppName)
	}

	consumer.Consume(ctx, event.DetailType, string(event.Detail))

	return nil
}
