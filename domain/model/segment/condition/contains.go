package condition

import "github.com/Luis-Miguel-BL/tiamat-notification/domain/model/segment"

const SegmentConditionTypeContains segment.ConditionType = "contains"

type Contains struct {
	*segment.ConditionBase
}

func (c *Contains) IsMatch(customer *segment.SerializedCustomer) (match bool, err error) {
	return
}
