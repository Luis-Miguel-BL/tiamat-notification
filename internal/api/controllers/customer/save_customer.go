package controller

import (
	"context"
	"net/http"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api/controllers"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/errors"
	usecase "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/ignition"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/ignition/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
)

type Request struct {
	WorkspaceID      string         `json:"workspace_id"`
	ExternalID       string         `json:"external_id"`
	Name             string         `json:"name"`
	EmailAddress     string         `json:"email_address"`
	PhoneNumber      string         `json:"phone_number"`
	CustomAttributes map[string]any `json:"custom_attributes"`
}

func (r *Request) Validate() (err error) {
	if r.WorkspaceID == "" {
		return errors.NewInvalidEmptyParamError("workspace_id")
	}
	if r.ExternalID == "" {
		return errors.NewInvalidEmptyParamError("external_id")
	}
	if r.Name == "" {
		return errors.NewInvalidEmptyParamError("name")
	}
	if r.EmailAddress == "" && r.PhoneNumber == "" {
		return errors.NewInvalidEmptyParamError("email_address")
	}
	return nil
}

type SaveCustomerController struct {
	saveCustomerUsecase usecase.SaveCustomerUsecase
	log                 logger.Logger
}

func NewSaveCustomerController(saveCustomerUsecase usecase.SaveCustomerUsecase, log logger.Logger) *SaveCustomerController {
	return &SaveCustomerController{
		saveCustomerUsecase: saveCustomerUsecase,
		log:                 log,
	}
}

func (c *SaveCustomerController) Execute(ctx context.Context, req controllers.Request) (res controllers.Response) {
	request, err := controllers.ParseRequestBody[*Request](req)
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Body = err.Error()
		return
	}
	err = c.saveCustomerUsecase.SaveCustomer(ctx, input.SaveCustomerInput{
		WorkspaceID: request.WorkspaceID,
		ExternalID:  request.ExternalID,
		Name:        request.Name,
		Contact: input.Contact{
			EmailAddress: request.EmailAddress,
			PhoneNumber:  request.PhoneNumber,
		},
		CustomAttributes: request.CustomAttributes,
	})
	if err != nil {
		res.StatusCode = http.StatusBadRequest
	}
	res.StatusCode = http.StatusOK
	return res
}
