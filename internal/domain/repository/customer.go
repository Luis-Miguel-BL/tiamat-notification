package repository

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer model.Customer) (err error)
	GetByID(ctx context.Context, customerID model.CustomerID, workspaceID model.WorkspaceID) (customer *model.Customer, err error)
	GetByExternalID(ctx context.Context, externalID vo.ExternalID, workspaceID model.WorkspaceID) (customer *model.Customer, find bool, err error)
}
