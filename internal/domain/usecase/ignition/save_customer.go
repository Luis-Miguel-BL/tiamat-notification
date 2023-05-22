package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/ignition/command"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type SaveCustomerUsecase struct {
	repo repository.CustomerRepository
}

func NewSaveCustomerUsecase(repo repository.CustomerRepository) *SaveCustomerUsecase {
	return &SaveCustomerUsecase{
		repo: repo,
	}
}

func (uc *SaveCustomerUsecase) SaveCustomer(ctx context.Context, command command.SaveCustomerCommand) (err error) {
	err = command.Validate()
	if err != nil {
		return err
	}
	customerID := model.CustomerID(command.CustomerID)
	workspaceID := model.WorkspaceID(command.WorkspaceID)

	customerName, err := vo.NewPersonName(command.Name)
	if err != nil {
		return err
	}

	contact, err := vo.NewContact(command.Contact.EmailAddress, command.Contact.PhoneNumber)
	if err != nil {
		return err
	}

	customAttr, err := vo.NewCustomAttributes(command.CustomAttributes)
	if err != nil {
		return err
	}

	customerToCreate, err := model.NewCustomer(
		model.NewCustomerInput{
			CustomerID:       customerID,
			WorkspaceID:      workspaceID,
			Name:             customerName,
			Contact:          contact,
			CustomAttributes: customAttr,
		},
	)
	if err != nil {
		return err
	}

	err = uc.repo.Save(ctx, *customerToCreate)

	return err
}
