package controllers

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api/controllers/request"
	usecase "github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/ignition"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/ignition/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
)

type CustomerController struct {
	saveCustomerUsecase        usecase.SaveCustomerUsecase
	createCustomerEventUsecase usecase.CreateCustomerEventUsecase
	log                        logger.Logger
}

func NewCustomerController(saveCustomerUsecase usecase.SaveCustomerUsecase, createCustomerEventUsecase usecase.CreateCustomerEventUsecase, log logger.Logger) *CustomerController {
	return &CustomerController{
		saveCustomerUsecase:        saveCustomerUsecase,
		createCustomerEventUsecase: createCustomerEventUsecase,
		log:                        log,
	}
}

func (c *CustomerController) SaveCustomer(ctx context.Context, request request.SaveCustomer) (err error) {
	err = request.Validate()
	if err != nil {
		return err
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
		return err
	}

	return nil
}
