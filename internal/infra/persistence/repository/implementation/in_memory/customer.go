package in_memory

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging"
)

type InMemoryCustomerRepo struct {
	data       map[model.WorkspaceID]map[model.CustomerID]model.Customer
	dispatcher messaging.AggregateEventDispatcher
}

func NewInMemoryCustomerRepo(dispatcher messaging.AggregateEventDispatcher) InMemoryCustomerRepo {
	return InMemoryCustomerRepo{
		dispatcher: dispatcher,
		data:       make(map[model.WorkspaceID]map[model.CustomerID]model.Customer),
	}
}

func (r InMemoryCustomerRepo) Save(ctx context.Context, customer model.Customer) (err error) {
	_, ok := r.data[customer.WorkspaceID()]
	if !ok {
		r.data[customer.WorkspaceID()] = make(map[model.CustomerID]model.Customer)
	}

	r.data[customer.WorkspaceID()][customer.CustomerID()] = customer

	err = r.dispatcher.PublishUncommitedEvents(ctx, *customer.AggregateRoot)
	if err != nil {
		return err
	}

	return nil
}

func (r InMemoryCustomerRepo) GetByID(ctx context.Context, customerID model.CustomerID, workspaceID model.WorkspaceID) (customer *model.Customer, err error) {
	customers, ok := r.data[workspaceID]
	if !ok {
		return customer, domain.NewEntityNotFoundError("customer")
	}
	dataCustomer, ok := customers[customerID]
	if !ok {
		return customer, domain.NewEntityNotFoundError("customer")
	}

	customer = &dataCustomer

	return customer, nil
}

func (r InMemoryCustomerRepo) GetByExternalID(ctx context.Context, externalID vo.ExternalID, workspaceID model.WorkspaceID) (customer *model.Customer, find bool, err error) {
	customers, ok := r.data[workspaceID]
	if !ok {
		return customer, false, nil
	}
	for _, dataCustomer := range customers {
		if dataCustomer.ExternalID() == externalID {
			return &dataCustomer, true, nil
		}
	}

	return customer, false, nil
}
