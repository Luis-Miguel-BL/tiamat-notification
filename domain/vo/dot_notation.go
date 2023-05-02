package vo

import (
	"fmt"

	"github.com/Luis-Miguel-BL/tiamat-notification/domain/util"
)

type DotNotation string

func (vo DotNotation) String() string {
	return string(vo)
}

func (vo *DotNotation) Validate() error {
	if util.IsEmpty(vo.String()) {
		return fmt.Errorf("dot_notation is empty")
	}
	return nil
}
