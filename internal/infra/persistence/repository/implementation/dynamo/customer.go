package dynamo

import (
	"context"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	dynamo_model "github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/persistence/model"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	customerTableName = "customer"
)

type DynamoCustomerRepo struct {
	client *DynamoClient
}

func (r *DynamoCustomerRepo) Save(ctx context.Context, customer model.Customer) (err error) {
	dynamoModel := dynamo_model.DynamoCustomer{}
	dynamoModel.ToRepo(customer)

	items := []map[string]types.AttributeValue{}
	items = append(items, dynamoModel.Customer)
	items = append(items, dynamoModel.Events...)
	items = append(items, dynamoModel.Segments...)

	err = r.client.BatchWrite(ctx, items, customerTableName)

	return err
}
func (r *DynamoCustomerRepo) GetByID(ctx context.Context, customerID model.CustomerID, workspaceID model.WorkspaceID) (customer model.Customer, err error) {
	return
}
func (r *DynamoCustomerRepo) GetStepJourney(ctx context.Context, workspaceID model.WorkspaceID, stepJourneyID model.StepJourneyID) (stepJourney model.StepJourney, err error) {
	return
}
func (r *DynamoCustomerRepo) SaveStepJourney(ctx context.Context, stepJourney model.StepJourney) (err error) {
	return
}
