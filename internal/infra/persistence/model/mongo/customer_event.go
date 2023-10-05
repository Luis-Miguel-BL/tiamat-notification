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
	CustomerEventID  string         `bson:"customer_event_id"`
	CustomerID       string         `bson:"customer_id"`
	WorkspaceID      string         `bson:"workspace_id"`
	Slug             string         `bson:"slug"`
	CustomAttributes map[string]any `bson:"custom_attributes"`
	OccurredAt       uint32         `bson:"occured_at"`
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
