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

type SaveCustomerController struct {
	uc  *usecase.SaveCustomerUsecase
	log logger.Logger
}

func NewSaveCustomerController(saveCustomerUsecase *usecase.SaveCustomerUsecase, log logger.Logger) *SaveCustomerController {
	return &SaveCustomerController{
		uc:  saveCustomerUsecase,
		log: log,
	}
}

func (c *SaveCustomerController) Execute(ctx context.Context, rawRequest api.Request) (res api.Response) {
	req, err := marshaller.UnmarshalRequestBody[*request.SaveCustomer](rawRequest)
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Body = err.Error()
		return
	}

	err = c.uc.SaveCustomer(ctx, input.SaveCustomerInput{
		WorkspaceID: req.WorkspaceID,
		ExternalID:  req.ExternalID,
		Name:        req.Name,
		Contact: input.Contact{
			EmailAddress: req.EmailAddress,
			PhoneNumber:  req.PhoneNumber,
		},
		CustomAttributes: req.CustomAttributes,
	})
	if err != nil {
		res.StatusCode = http.StatusBadRequest
	}

	res.StatusCode = http.StatusOK

	return res
}
