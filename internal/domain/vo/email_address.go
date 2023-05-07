package vo

import (
	"regexp"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type EmailAddress string

var regexEmail = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func NewEmailAddress(email string) (emailAddress EmailAddress, err domain.DomainError) {
	if util.IsEmpty(email) {
		return emailAddress, domain.NewInvalidEmptyParamError("email")
	}

	regex, err := regexp.Compile(regexEmail)
	if err != nil {
		return emailAddress, domain.NewInvalidEmptyParamError("email")
	}

	if !regex.MatchString(email) {
		return emailAddress, domain.NewInvalidEmptyParamError("email")
	}
	return EmailAddress(email), nil
}

func (vo EmailAddress) String() string {
	return string(vo)
}
