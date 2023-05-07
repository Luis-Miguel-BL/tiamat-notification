package repository

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer model.Customer) (err error)
	CreateCustomerEvent(ctx context.Context, customerID model.CustomerID, workspaceID model.WorkspaceID, customerEvent model.CustomerEvent) (err error)
	GetByID(ctx context.Context, customerID model.CustomerID, workspaceID model.WorkspaceID) (customer model.Customer, err error)
}
