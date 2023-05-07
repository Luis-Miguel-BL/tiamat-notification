package vo

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type DotNotation string

func NewDotNotation(str string) (dotNotation DotNotation, err domain.DomainError) {

	if util.IsEmpty(str) {
		return dotNotation, domain.NewInvalidEmptyParamError("dot-notation")
	}
	return DotNotation(str), nil
}

func (vo DotNotation) String() string {
	return string(vo)
}
