package usecase

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase/ignition/input"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/repository"
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
	email, err := vo.NewEmailAddress(input.Contact.EmailAddress)
	if err != nil {
		return err
	}
	phone, err := vo.NewPhoneNumber(input.Contact.PhoneNumber)
	if err != nil {
		return err
	}
	customAttr, err := vo.NewCustomAttributes(input.CustomAttributes)
	if err != nil {
		return err
	}
	customer, find, err := uc.repo.GetByExternalID(ctx, externalID, workspaceID)
	if err != nil {
		return err
	}
	if !find {
		contact, err := vo.NewContact(input.Contact.EmailAddress, input.Contact.PhoneNumber)
		if err != nil {
			return err
		}
		customer, err = model.NewCustomer(
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
	} else {
		customer.Update(model.UpdateCustomerInput{
			Name:       customerName,
			Email:      email,
			Phone:      phone,
			CustomAttr: customAttr,
		})
	}
	err = uc.repo.Save(ctx, *customer)

	return err
}
