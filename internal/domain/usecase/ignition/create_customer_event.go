package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/ignition/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type CreateCustomerEventUsecase struct {
	repo repository.CustomerRepository
}

func NewCreateCustomerEventUsecase(repo repository.CustomerRepository) *CreateCustomerEventUsecase {
	return &CreateCustomerEventUsecase{
		repo: repo,
	}
}

func (uc *CreateCustomerEventUsecase) CreateCustomerEvent(ctx context.Context, command command.CreateCustomerEventCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	customerID := model.CustomerID(command.CustomerID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)

	eventSlug, err := vo.NewSlug(command.Slug)
	if err != nil {
		return err
	}
	customAttr, err := vo.NewCustomAttributes(command.CustomAttributes)
	if err != nil {
		return err
	}

	customerEventToCreate, err := model.NewCustomerEvent(
		model.NewCustomerEventInput{
			Slug:             eventSlug,
			CustomAttributes: customAttr,
		},
	)
	if err != nil {
		return err
	}

	err = uc.repo.CreateCustomerEvent(ctx, customerID, workspaceID, *customerEventToCreate)

	return err
}
