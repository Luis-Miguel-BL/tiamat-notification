package model

import (
	"fmt"
	"strconv"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
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

func (m *DynamoCustomerEvent) ToDomain(item map[string]types.AttributeValue) (customer model.Customer) {
	attributevalue.UnmarshalMap(item, m)

	return
}
func (m *DynamoCustomer) ToRepo(customer model.Customer) (err error) {
	customerPK := MakeCustomerPK(string(customer.WorkspaceID()), string(customer.CustomerID()))

	m.Customer, err = attributevalue.MarshalMap(makeCustomerMap(customer, customerPK))
	if err != nil {
		return err
	}

	for _, customerEvents := range customer.Events() {
		for _, customerEvent := range customerEvents {
			dynamoCustomerEvent, err := attributevalue.MarshalMap(makeCustomerEventMap(customerEvent, customerPK))
			if err != nil {
				return err
			}
			m.Events = append(m.Events, dynamoCustomerEvent)
		}
	}
	for _, customerSegment := range customer.Segments() {
		dynamoCustomerSegment, err := attributevalue.MarshalMap(makeCustomerSegmentMap(customerSegment, customerPK))
		if err != nil {
			return err
		}
		m.Segments = append(m.Segments, dynamoCustomerSegment)
	}

	return nil

}

func makeCustomerMap(customer model.Customer, customerPK string) (customerMap map[string]interface{}) {
	contact := customer.Contact()
	customerMap = map[string]interface{}{
		"customer_id":  string(customer.CustomerID()),
		"workspace_id": string(customer.WorkspaceID()),
		"external_id":  customer.ExternalID().String(),
		"name":         customer.Name().String(),
		"contact": map[string]interface{}{
			"email": map[string]interface{}{
				"email_address":   contact.Email.EmailAddress.String(),
				"unsubscribed_at": strconv.FormatUint(uint64(contact.Email.UnsubscribedAt.Unix()), 10),
				"bounced_at":      strconv.FormatUint(uint64(contact.Email.BouncedAt.Unix()), 10),
			},
			"phone": map[string]interface{}{
				"phone_number":             contact.Phone.PhoneNumber.String(),
				"sms_unsubscribed_at":      strconv.FormatUint(uint64(contact.Phone.SMSUnsubscribedAt.Unix()), 10),
				"whatsapp_unsubscribed_at": strconv.FormatUint(uint64(contact.Phone.WhatsAppUnsubscribedAt.Unix()), 10),
			},
		},
		"custom_attributes": customer.CustomAttributes(),
		"PK":                customerPK,
		"SK":                makeCustomerSK(string(customer.WorkspaceID()), customer.ExternalID().String()),
	}

	return customerMap
}

func makeCustomerEventMap(customerEvent model.CustomerEvent, customerPK string) (customerEventMap map[string]interface{}) {
	customerEventMap = map[string]interface{}{
		"customer_event_id": string(customerEvent.CustomerEventID()),
		"customer_id":       string(customerEvent.CustomerID()),
		"workspace_id":      string(customerEvent.WorkspaceID()),
		"slug":              customerEvent.Slug().String(),
		"custom_attributes": customerEvent.CustomAttributes(),
		"occurred_at":       strconv.FormatUint(uint64(customerEvent.OccurredAt().Unix()), 10),
		"PK":                customerPK,
		"SK":                makeCustomerEventSK(string(customerEvent.WorkspaceID()), string(customerEvent.CustomerEventID())),
	}

	return customerEventMap
}

func makeCustomerSegmentMap(customerSegment model.CustomerSegment, customerPK string) (customerSegmentMap map[string]interface{}) {
	customerSegmentMap = map[string]interface{}{
		"customer_id":  string(customerSegment.CustomerID()),
		"workspace_id": string(customerSegment.WorkspaceID()),
		"segment_id":   string(customerSegment.SegmentID()),
		"matched_at":   strconv.FormatUint(uint64(customerSegment.MatchedAt().Unix()), 10),
		"PK":           customerPK,
		"SK":           makeCustomerSegmentSK(string(customerSegment.WorkspaceID()), string(customerSegment.SegmentID())),
	}

	return customerSegmentMap
}

func MakeCustomerPK(workspaceID string, customerID string) (pk string) {
	return fmt.Sprintf("CUSTOMER#%s#%s", workspaceID, customerID)
}

const customerSKPrefix = "#CUSTOMER"

func makeCustomerSK(workspaceID string, externalID string) (sk string) {
	return fmt.Sprintf("%s#%s#%s", customerSKPrefix, workspaceID, externalID)
}

const customerEventSKPrefix = "CUSTOMER_EVENT"

func makeCustomerEventSK(workspaceID string, eventID string) (sk string) {
	return fmt.Sprintf("%s#%s#%s", customerEventSKPrefix, workspaceID, eventID)
}

const customerSegmentSKPrefix = "CUSTOMER_SEGMENT"

func makeCustomerSegmentSK(workspaceID string, segmentID string) (sk string) {
	return fmt.Sprintf("%s#%s#%s", customerSegmentSKPrefix, workspaceID, segmentID)
}
