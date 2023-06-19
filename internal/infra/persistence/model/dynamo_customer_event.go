package model

import (
	"fmt"
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoCustomerEvent struct {
	PK               string         `dynamodbav:"PK"`
	SK               string         `dynamodbav:"SK"`
	CustomerEventID  string         `dynamodbav:"customer_event_id"`
	CustomerID       string         `dynamodbav:"customer_id"`
	WorkspaceID      string         `dynamodbav:"workspace_id"`
	Slug             string         `dynamodbav:"slug"`
	CustomAttributes map[string]any `dynamodbav:"custom_attributes"`
	OccurredAt       uint32         `dynamodbav:"occured_at"`
}

func (m *DynamoCustomerEvent) ToDomain(item map[string]types.AttributeValue) (customerEvent *model.CustomerEvent, err error) {
	err = attributevalue.UnmarshalMap(item, m)
	if err != nil {
		return customerEvent, err
	}

	eventSlug, err := vo.NewSlug(m.Slug)
	if err != nil {
		return customerEvent, err
	}
	customAttr, err := vo.NewCustomAttributes(m.CustomAttributes)
	if err != nil {
		return customerEvent, err
	}
	customerEvent, err = model.NewCustomerEventToRepo(model.NewCustomerEventToRepoInput{
		CustomerEventID:  model.CustomerEventID(m.CustomerEventID),
		CustomerID:       model.CustomerID(m.CustomerID),
		WorkspaceID:      model.WorkspaceID(m.WorkspaceID),
		Slug:             eventSlug,
		CustomAttributes: customAttr,
		OccurredAt:       time.Unix(int64(m.OccurredAt), 0),
	})

	return customerEvent, nil
}

func (m *DynamoCustomerEvent) ToRepo(customerEvent model.CustomerEvent) (item map[string]types.AttributeValue, err error) {
	customerPK := MakeCustomerPK(string(customerEvent.WorkspaceID()), string(customerEvent.CustomerID()))
	customerEventSK := makeCustomerEventSK(string(customerEvent.WorkspaceID()), string(customerEvent.CustomerEventID()))

	m.PK = customerPK
	m.SK = customerEventSK
	m.CustomerEventID = string(customerEvent.CustomerEventID())
	m.CustomerID = string(customerEvent.CustomerID())
	m.WorkspaceID = string(customerEvent.WorkspaceID())
	m.Slug = customerEvent.Slug().String()
	m.CustomAttributes = customerEvent.CustomAttributes()
	m.OccurredAt = uint32(customerEvent.OccurredAt().Unix())

	item, err = attributevalue.MarshalMap(m)
	if err != nil {
		return item, err
	}
	return item, nil

}

const customerEventSKPrefix = "CUSTOMER_EVENT"

func makeCustomerEventSK(workspaceID string, eventID string) (sk string) {
	return fmt.Sprintf("%s#%s#%s", customerEventSKPrefix, workspaceID, eventID)
}
