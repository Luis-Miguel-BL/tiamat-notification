package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/usecase/ignition/input"
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

func (uc *SaveCustomerUsecase) SaveCustomer(ctx context.Context, input input.SaveCustomerInput) (err error) {
	err = input.Validate()
	if err != nil {
		return err
	}
	workspaceID := model.WorkspaceID(input.WorkspaceID)

	externalID, err := vo.NewExternalID(input.ExternalID)
	if err != nil {
		return err
	}
	customerName, err := vo.NewPersonName(input.Name)
	if err != nil {
		return err
	}

	contact, err := vo.NewContact(input.Contact.EmailAddress, input.Contact.PhoneNumber)
	if err != nil {
		return err
	}

	customAttr, err := vo.NewCustomAttributes(input.CustomAttributes)
	if err != nil {
		return err
	}

	customerToCreate, err := model.NewCustomer(
		model.NewCustomerInput{
			ExternalID:       externalID,
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
