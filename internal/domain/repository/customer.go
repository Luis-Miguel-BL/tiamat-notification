package repository

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer model.Customer) (err error)
	CreateCustomerEvent(ctx context.Context, customerID model.CustomerID, workspaceID model.WorkspaceID, customerEvent model.CustomerEvent) (err error)
	GetByID(ctx context.Context, customerID model.CustomerID, workspaceID model.WorkspaceID) (customer model.Customer, err error)
	GetStepJourney(ctx context.Context, workspaceID model.WorkspaceID, stepJourneyID model.StepJourneyID) (stepJourney model.StepJourney, err error)
	SaveStepJourney(ctx context.Context, stepJourney model.StepJourney) (err error)
}
