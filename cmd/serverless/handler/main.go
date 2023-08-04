package handler

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/aws/aws-lambda-go/events"
)

type ServerlessHandler struct {
	cfg            *config.Config
	log            logger.Logger
	usecaseManager *usecase.UsecaseManager
}

func NewServerlessHandler(cfg *config.Config, log logger.Logger, usecaseManager *usecase.UsecaseManager) *ServerlessHandler {
	return &ServerlessHandler{
		cfg:            cfg,
		log:            log,
		usecaseManager: usecaseManager,
	}
}

func (h *ServerlessHandler) Handler(ctx context.Context, e any) (response any, err error) {
	switch event := e.(type) {
	case events.SQSEvent:
		response, err = h.sqsHandle(ctx, event)
	case events.APIGatewayProxyRequest:
		response, err = h.apiGatewayHandle(ctx, event)
	}
	return response, err
}
