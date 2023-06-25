package model

import (
	"fmt"
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/vo"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoCustomer struct {
	PK               string                  `dynamodbav:"PK"`
	SK               string                  `dynamodbav:"SK"`
	CustomerID       string                  `dynamodbav:"customer_id"`
	WorkspaceID      string                  `dynamodbav:"workspace_id"`
	ExternalID       string                  `dynamodbav:"external_id"`
	Name             string                  `dynamodbav:"name"`
	Contact          DynamoCustomerContact   `dynamodbav:"name"`
	CustomAttributes map[string]any          `dynamodbav:"name"`
	CreatedAt        uint32                  `dynamodbav:"created_at"`
	UpdatedAt        uint32                  `dynamodbav:"updated_at"`
	Events           []DynamoCustomerEvent   `dynamodbav:"-"`
	Segments         []DynamoCustomerSegment `dynamodbav:"-"`
}
type DynamoCustomerContact struct {
	Email struct {
		EmailAddress   string `dynamodbav:"email_address"`
		UnsubscribedAt uint32 `dynamodbav:"unsubscribed_at"`
		BouncedAt      uint32 `dynamodbav:"bounced_at"`
	} `dynamodbav:"email"`
	Phone struct {
		PhoneNumber            string `dynamodbav:"phone_number"`
		SMSUnsubscribedAt      uint32 `dynamodbav:"sms_unsubscribed_at"`
		WhatsAppUnsubscribedAt uint32 `dynamodbav:"whatsapp_unsubscribed_at"`
	} `dynamodbav:"phone"`
}

func (m *DynamoCustomer) ToDomain(items []map[string]types.AttributeValue) (customer *model.Customer, err error) {
	customerEvents := make(map[vo.Slug][]model.CustomerEvent)
	customerSegments := make(map[model.SegmentID]model.CustomerSegment)
	for _, item := range items {
		var sk string
		err := attributevalue.Unmarshal(item["SK"], &sk)
		if err != nil {
			return customer, err
		}
		switch true {
		case strings.HasPrefix(sk, customerSKPrefix):
			attributevalue.UnmarshalMap(item, m)
		case strings.HasPrefix(sk, customerEventSKPrefix):
			dynamoCustomerEvent := DynamoCustomerEvent{}
			customerEvent, err := dynamoCustomerEvent.ToDomain(item)
			if err != nil {
				return customer, err
			}
			m.Events = append(m.Events, dynamoCustomerEvent)
			customerEvents[customerEvent.Slug()] = append(customerEvents[customerEvent.Slug()], *customerEvent)
		case strings.HasPrefix(sk, customerSegmentSKPrefix):
			dynamoCustomerSegment := DynamoCustomerSegment{}
			customerSegment, err := dynamoCustomerSegment.ToDomain(item)
			if err != nil {
				return customer, err
			}
			m.Segments = append(m.Segments, dynamoCustomerSegment)
			customerSegments[customerSegment.SegmentID()] = *customerSegment
		}
	}
	externalID, err := vo.NewExternalID(m.ExternalID)
	if err != nil {
		return customer, err
	}
	name, err := vo.NewPersonName(m.Name)
	if err != nil {
		return customer, err
	}
	contact, err := vo.NewContactToRepo(vo.NewContactToRepoInput{
		EmailAddress:           m.Contact.Email.EmailAddress,
		UnsubscribedAt:         util.NewUnixTime(m.Contact.Email.UnsubscribedAt),
		BouncedAt:              util.NewUnixTime(m.Contact.Email.BouncedAt),
		PhoneNumber:            m.Contact.Phone.PhoneNumber,
		SMSUnsubscribedAt:      util.NewUnixTime(m.Contact.Phone.SMSUnsubscribedAt),
		WhatsAppUnsubscribedAt: util.NewUnixTime(m.Contact.Phone.WhatsAppUnsubscribedAt),
	})
	if err != nil {
		return customer, err
	}
	customer, err = model.RestoreToRepo(model.RestoreToRepoInput{
		CustomerID:       model.CustomerID(m.CustomerID),
		WorkspaceID:      model.WorkspaceID(m.WorkspaceID),
		ExternalID:       externalID,
		Name:             name,
		Contact:          contact,
		CustomAttributes: m.CustomAttributes,
		Events:           customerEvents,
		Segments:         customerSegments,
		CreatedAt:        util.NewUnixTime(m.CreatedAt),
		UpdatedAt:        util.NewUnixTime(m.UpdatedAt),
	})
	if err != nil {
		return customer, err
	}

	return customer, nil
}
func (m *DynamoCustomer) ToRepo(customer model.Customer) (items []map[string]types.AttributeValue, err error) {
	items = make([]map[string]types.AttributeValue, 0)
	customerPK := MakeCustomerPK(string(customer.WorkspaceID()), string(customer.CustomerID()))
	customerSK := MakeCustomerSK(string(customer.WorkspaceID()), string(customer.ExternalID()))

	m.PK = customerPK
	m.SK = customerSK

	m.CustomerID = string(customer.CustomerID())
	m.WorkspaceID = string(customer.WorkspaceID())
	m.ExternalID = string(customer.ExternalID())
	m.Name = customer.Name().String()

	customerContact := customer.Contact()
	m.Contact.Email.EmailAddress = customerContact.Email.EmailAddress.String()
	m.Contact.Email.UnsubscribedAt = uint32(customerContact.Email.UnsubscribedAt.Unix())
	m.Contact.Email.BouncedAt = uint32(customerContact.Email.BouncedAt.Unix())
	m.Contact.Phone.PhoneNumber = customerContact.Phone.PhoneNumber.String()
	m.Contact.Phone.SMSUnsubscribedAt = uint32(customerContact.Phone.SMSUnsubscribedAt.Unix())
	m.Contact.Phone.WhatsAppUnsubscribedAt = uint32(customerContact.Phone.WhatsAppUnsubscribedAt.Unix())

	m.CustomAttributes = customer.CustomAttributes()
	m.CreatedAt = uint32(customer.CreatedAt().Unix())
	m.UpdatedAt = uint32(customer.UpdatedAt().Unix())

	dynamoCustomerEvents := make([]DynamoCustomerEvent, 0)
	for _, events := range customer.Events() {
		for _, event := range events {
			dynamoCustomerEvent := DynamoCustomerEvent{}
			item, _ := dynamoCustomerEvent.ToRepo(event)
			dynamoCustomerEvents = append(dynamoCustomerEvents, dynamoCustomerEvent)
			items = append(items, item)
		}
	}
	m.Events = dynamoCustomerEvents

	dynamoCustomerSegments := make([]DynamoCustomerSegment, 0)
	for _, segment := range customer.Segments() {
		dynamoCustomerSegment := DynamoCustomerSegment{}
		item, _ := dynamoCustomerSegment.ToRepo(segment)
		dynamoCustomerSegments = append(dynamoCustomerSegments, dynamoCustomerSegment)
		items = append(items, item)
	}
	m.Segments = dynamoCustomerSegments

	item, err := attributevalue.MarshalMap(m)
	if err != nil {
		return items, err
	}
	items = append(items, item)

	return items, nil
}

func MakeCustomerPK(workspaceID string, customerID string) (pk string) {
	return fmt.Sprintf("CUSTOMER#%s#%s", workspaceID, customerID)
}

const customerSKPrefix = "#CUSTOMER"

func MakeCustomerSK(workspaceID string, externalID string) (sk string) {
	return fmt.Sprintf("%s#%s#%s", customerSKPrefix, workspaceID, externalID)
}
