package customer

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/model/campaign"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/model/segment"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/model/workspace"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/util"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/vo"
)

var AggregateType = domain.AggregateType("customer")

type CustomerID string
type Customer struct {
	*domain.Aggregate
	CustomerID      CustomerID
	WorkspaceID     workspace.WorkspaceID
	Name            vo.CustomerName
	Contact         vo.Contact
	Events          map[vo.Slug][]CustomerEvent
	CurrentActions  map[campaign.CampaignID]CurrentAction
	CurrentSegments map[segment.SegmentID]CurrentSegment
	CustomData      map[string]interface{}
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewCustomer(customerID string, workspaceID string, name string) (customer *Customer) {
	return &Customer{
		CustomerID:  CustomerID(customerID),
		WorkspaceID: workspace.WorkspaceID(workspaceID),
		Name:        vo.CustomerName(name),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (e *Customer) Validate() error {
	if util.IsEmpty(string(e.CustomerID)) {
		return domain.NewInvalidEmptyParamError("customer_id")
	}
	if util.IsEmpty(string(e.WorkspaceID)) {
		return domain.NewInvalidEmptyParamError("workspace_id")
	}
	if err := e.Name.Validate(); err != nil {
		return err
	}
	if err := e.Contact.Validate(); err != nil {
		return err
	}
	return nil
}

func (e *Customer) AddContact(contact vo.Contact) error {
	if err := contact.Validate(); err != nil {
		return err
	}
	e.Contact = contact
	return nil
}
