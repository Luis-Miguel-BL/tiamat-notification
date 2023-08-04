package controller

import (
	"context"
	"net/http"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api/ignition/request"
	usecase "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/ignition"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/ignition/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging/marshaller"
)

type CreateCustomerEventController struct {
	uc  usecase.CreateCustomerEventUsecase
	log logger.Logger
}

func NewCreateCustomerEventController(saveCustomerUsecase usecase.CreateCustomerEventUsecase, log logger.Logger) *CreateCustomerEventController {
	return &CreateCustomerEventController{
		uc:  saveCustomerUsecase,
		log: log,
	}
}

func (c *CreateCustomerEventController) Execute(ctx context.Context, rawRequest api.Request) (res api.Response) {
	req, err := marshaller.UnmarshalRequestBody[*request.CreateCustomerEvent](rawRequest)
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Body = err.Error()
		return
	}

	err = c.uc.CreateCustomerEvent(ctx, input.CreateCustomerEventInput{
		WorkspaceID:      req.WorkspaceID,
		CustomerID:       req.CustomerID,
		CustomerEventID:  req.CustomerEventID,
		Slug:             req.Slug,
		CustomAttributes: req.CustomAttributes,
	})
	if err != nil {
		res.StatusCode = http.StatusBadRequest
	}

	res.StatusCode = http.StatusOK
	return res
}
