package handler

import (
	"context"
	"net/http"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api"
	"github.com/aws/aws-lambda-go/events"
)

func (h *ServerlessHandler) apiGatewayHandle(ctx context.Context, request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	var controllerFunc api.ControllerFunc
	switch request.Path {
	case "identify":
		controller = customer_controller.NewSaveCustomerController(h.usecaseManager.SaveCustomer, h.log)
	default:
		response.StatusCode = http.StatusNotFound
		return
	}
	res := controllerFunc(ctx, api.Request{
		Method: request.HTTPMethod,
		Body:   request.Body,
	})

	response.Body = res.Body
	response.StatusCode = res.StatusCode

	return response, nil
}
