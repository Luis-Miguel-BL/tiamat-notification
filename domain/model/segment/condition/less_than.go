package condition

import "github.com/Luis-Miguel-BL/tiamat-notification/domain/model/segment"

const SegmentConditionTypeLessThan segment.ConditionType = "contains"

type LessThan struct {
	*segment.ConditionBase
}

func (c *LessThan) IsMatch(customer *segment.SerializedCustomer) (match bool, err error) {
	return
}
