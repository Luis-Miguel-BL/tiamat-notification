package handler

import (
	"context"
	"net/http"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api/controllers"
	customer_controller "github.com/Luis-Miguel-BL/tiamat-notification/internal/api/controllers/customer"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/registry"
	"github.com/aws/aws-lambda-go/events"
)

func ApiGatewayHandle(ctx context.Context, request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	registry := ctx.Value("registry").(*registry.Registry)

	log := registry.Inject("logger").(*logger.ZeroLogger)
	usecases := registry.Inject("usecases").(*usecase.UsecaseManager)

	var controller controllers.Controller
	switch request.Path {
	case "identify":
		controller = customer_controller.NewSaveCustomerController(*usecases.SaveCustomer, log)
	default:
		response.StatusCode = http.StatusNotFound
		return
	}
	res := controller.Execute(ctx, controllers.Request{
		Method: request.HTTPMethod,
		Body:   request.Body,
	})

	response.Body = res.Body
	response.StatusCode = res.StatusCode

	return response, nil
}
