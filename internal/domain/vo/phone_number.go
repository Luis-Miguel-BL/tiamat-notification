package vo

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type PhoneNumber string

func NewPhoneNumber(phone string) (phoneNumber PhoneNumber, err domain.DomainError) {
	if util.IsEmpty(phone) {
		return phoneNumber, domain.NewInvalidEmptyParamError("phone_number")
	}
	//@todo implements regex to validate phone_number
	return PhoneNumber(phone), nil
}

func (vo PhoneNumber) String() string {
	return string(vo)
}
