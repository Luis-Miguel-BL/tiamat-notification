package dynamo

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/messaging"
	dynamo_model "github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/persistence/model/dynamo"
)

const (
	customerTableName    = "customer"
	customerTableSKIndex = "customer-sk-index"
)

type DynamoCustomerRepo struct {
	client     *DynamoClient
	dispatcher messaging.AggregateEventDispatcher
}

func NewDynamoCustomerRepo(client *DynamoClient, dispatcher messaging.AggregateEventDispatcher) DynamoCustomerRepo {
	return DynamoCustomerRepo{
		client:     client,
		dispatcher: dispatcher,
	}
}

func (r DynamoCustomerRepo) Save(ctx context.Context, customer model.Customer) (err error) {
	dynamoModel := dynamo_model.DynamoCustomer{}
	items, err := dynamoModel.ToRepo(customer)
	if err != nil {
		return err
	}
	err = r.client.BatchWrite(ctx, items, customerTableName)
	if err != nil {
		return err
	}

	err = r.dispatcher.PublishUncommitedEvents(ctx, *customer.AggregateRoot)
	if err != nil {
		return err
	}

	return nil
}

func (r DynamoCustomerRepo) GetByID(ctx context.Context, customerID model.CustomerID, workspaceID model.WorkspaceID) (customer *model.Customer, err error) {
	dynamoCustomer := dynamo_model.DynamoCustomer{}
	dynamoResult, count, err := r.client.QueryByPK(
		ctx,
		customerTableName,
		dynamo_model.MakeCustomerPK(string(workspaceID), string(customerID)),
	)
	if err != nil {
		return customer, err
	}
	if count == 0 {
		return customer, domain.NewEntityNotFoundError("customer")
	}

	customer, err = dynamoCustomer.ToDomain(dynamoResult)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r DynamoCustomerRepo) GetByExternalID(ctx context.Context, externalID vo.ExternalID, workspaceID model.WorkspaceID) (customer *model.Customer, find bool, err error) {
	dynamoCustomer := dynamo_model.DynamoCustomer{}
	dynamoResult, count, err := r.client.QueryByIndex(
		ctx,
		customerTableName,
		customerTableSKIndex,
		"SK",
		dynamo_model.MakeCustomerSK(string(workspaceID), string(externalID)),
	)
	if err != nil {
		return customer, false, err
	}
	if count == 0 {
		return customer, false, domain.NewEntityNotFoundError("customer")
	}

	customer, err = dynamoCustomer.ToDomain(dynamoResult)
	if err != nil {
		return customer, false, err
	}

	return customer, true, nil
}
