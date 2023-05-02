package condition

import "github.com/Luis-Miguel-BL/tiamat-notification/domain/model/segment"

const SegmentConditionTypeEqual segment.ConditionType = "contains"

type Equal struct {
	*segment.ConditionBase
}

func (c *Equal) IsMatch(customer *segment.SerializedCustomer) (match bool, err error) {
	return
}
