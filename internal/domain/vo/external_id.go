package vo

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type ExternalID string

func NewExternalID(externalID string) (newExternalID ExternalID, err domain.DomainError) {
	if util.IsEmpty(externalID) {
		return newExternalID, domain.NewInvalidEmptyParamError("external_id")
	}
	if len(externalID) > 250 {
		return newExternalID, domain.NewInvalidParamError("external_id")
	}
	return ExternalID(externalID), nil
}

func (vo ExternalID) String() string {
	return string(vo)
}
