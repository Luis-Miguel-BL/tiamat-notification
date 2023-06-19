package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoCustomer struct {
	PK               string         `dynamodbav:"PK"`
	SK               string         `dynamodbav:"SK"`
	CustomerID       string `dynamodbav:"customer_id"`
	WorkspaceID      string `dynamodbav:"workspace_id"`
	ExternalID       string `dynamodbav:"external_id"`
	Name             string `dynamodbav:"name"`
	Contact          vo.Contact
	CustomAttributes map[string]any  `dynamodbav:"name"`
	CreatedAt        uint32  `dynamodbav:"created_at"`
	UpdatedAt        uint32 `dynamodbav:"updated_at"`
	Events   []map[string]types.AttributeValue
	Segments []map[string]types.AttributeValue
}

func (m *DynamoCustomer) ToDomain(items []map[string]types.AttributeValue) (customer model.Customer) {
	for _, item := range items {
		sk, ok := item["SK"]
		sk.
		if  !ok {
			continue
		}
		if strings.Contains(sk, customerSKPrefix) {

		}
	}

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
