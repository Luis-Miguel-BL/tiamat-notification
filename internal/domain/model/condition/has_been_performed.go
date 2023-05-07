package condition

import "github.com/Luis-Miguel-BL/tiamat-notification/internal/domain/model"

const SegmentConditionTypeHasBeenPerformed model.ConditionType = "hass-been-performed"

type HasBeenPerformed struct {
	*model.ConditionBase
}

func (c *HasBeenPerformed) IsMatch(customer *model.SerializedCustomer) (match bool) {
	_, match = customer.Events[c.EventSlug]
	return match
}
