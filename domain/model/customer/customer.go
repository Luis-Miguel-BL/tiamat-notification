package customer

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/model/campaign"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/model/segment"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/model/workspace"
	"github.com/Luis-Miguel-BL/tiamat-notification/domain/vo"
)

var AggregateType = domain.AggregateType("customer")

type CustomerID string

func NewCustomerID(customerID string) CustomerID {
	return CustomerID(customerID)
}

type Customer struct {
	*domain.Aggregate
	CustomerID       CustomerID
	WorkspaceID      workspace.WorkspaceID
	Name             vo.PersonName
	Contact          vo.Contact
	CustomAttributes vo.CustomAttributes
	Events           map[vo.Slug][]CustomerEvent
	ActionsTriggered map[campaign.CampaignID][]ActionsTriggered
	CurrentSegments  map[segment.SegmentID]CurrentSegment
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type NewCustomerInput struct {
	CustomerID       CustomerID
	WorkspaceID      workspace.WorkspaceID
	Name             vo.PersonName
	Contact          vo.Contact
	CustomAttributes vo.CustomAttributes
}

func NewCustomer(input NewCustomerInput) (customer *Customer) {
	return &Customer{
		Aggregate:        domain.NewAggregate(AggregateType, domain.AggregateID(input.CustomerID)),
		CustomerID:       input.CustomerID,
		WorkspaceID:      input.WorkspaceID,
		Name:             input.Name,
		Contact:          input.Contact,
		CustomAttributes: input.CustomAttributes,
		Events:           make(map[vo.Slug][]CustomerEvent),
		ActionsTriggered: make(map[campaign.CampaignID][]ActionsTriggered),
		CurrentSegments:  make(map[segment.SegmentID]CurrentSegment),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func (e *Customer) Validate() error {
	if e.CustomerID == "" {
		return domain.NewInvalidEmptyParamError("customer_id")
	}
	if e.WorkspaceID == "" {
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

func (e *Customer) Serialize() (serialized segment.SerializedCustomer) {
	serialized.Attributes = segment.SerializedAttributes(e.CustomAttributes)
	serialized.Attributes["name"] = e.Name
	serialized.Attributes["first_name"] = e.Name.GetFirstName()
	serialized.Attributes["email"] = e.Contact.Email.EmailAddress
	serialized.Attributes["phone"] = e.Contact.Phone.PhoneNumber

	serialized.Events = make(map[vo.Slug]segment.SerializedAttributes)

}
func (e *Customer) AppendEvent(event CustomerEvent) {
	e.Events[event.Slug] = append(e.Events[event.Slug], event)
}

func (e *Customer) MatchWithSegment(segment segment.Segment) (matched bool) {
	matched = true
	for _, condition := range segment.Conditions {
		condition.IsMatch()
	}
}
