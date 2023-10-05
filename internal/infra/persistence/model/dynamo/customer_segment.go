package model

import (
	"fmt"
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoCustomerSegment struct {
	PK          string `dynamodbav:"PK"`
	SK          string `dynamodbav:"SK"`
	SegmentID   string `dynamodbav:"segment_id"`
	CustomerID  string `dynamodbav:"customer_id"`
	WorkspaceID string `dynamodbav:"workspace_id"`
	MatchedAt   uint32 `dynamodbav:"matched_at"`
}

func (m *DynamoCustomerSegment) ToDomain(item map[string]types.AttributeValue) (customerSegment *model.CustomerSegment, err error) {
	err = attributevalue.UnmarshalMap(item, m)
	if err != nil {
		return customerSegment, err
	}

	customerSegment, err = model.NewCustomerSegmentToRepo(model.NewCustomerSegmentToRepoInput{
		CustomerID:  model.CustomerID(m.CustomerID),
		WorkspaceID: model.WorkspaceID(m.WorkspaceID),
		SegmentID:   model.SegmentID(m.SegmentID),
		MatchedAt:   time.Unix(int64(m.MatchedAt), 0),
	})
	if err != nil {
		return customerSegment, err
	}

	return customerSegment, nil
}

func (m *DynamoCustomerSegment) ToRepo(customerSegment model.CustomerSegment) (item map[string]types.AttributeValue, err error) {
	customerPK := MakeCustomerPK(string(customerSegment.WorkspaceID()), string(customerSegment.CustomerID()))
	customerSegmentSK := makeCustomerSegmentSK(string(customerSegment.WorkspaceID()), string(customerSegment.SegmentID()))

	m.PK = customerPK
	m.SK = customerSegmentSK
	m.SegmentID = string(customerSegment.SegmentID())
	m.CustomerID = string(customerSegment.CustomerID())
	m.WorkspaceID = string(customerSegment.WorkspaceID())
	m.MatchedAt = uint32(customerSegment.MatchedAt().Unix())

	item, err = attributevalue.MarshalMap(m)
	if err != nil {
		return item, err
	}
	return item, nil

}

const customerSegmentSKPrefix = "CUSTOMER_SEGMENT"

func makeCustomerSegmentSK(workspaceID string, segmentID string) (sk string) {
	return fmt.Sprintf("%s#%s#%s", customerSegmentSKPrefix, workspaceID, segmentID)
}
