package vo

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/util"
)

type Contact struct {
	Email ContactEmail
	Phone ContactPhone
}

type ContactEmail struct {
	EmailAddress   EmailAddress
	UnsubscribedAt time.Time
	BouncedAt      time.Time
}

type ContactPhone struct {
	PhoneNumber            PhoneNumber
	SMSUnsubscribedAt      time.Time
	WhatsAppUnsubscribedAt time.Time
}

func NewContact(emailAddress string, phoneNumber string) (contact Contact, err error) {
	hasSomeContact := false
	if !util.IsEmpty(emailAddress) {
		hasSomeContact = true
		contact.Email.EmailAddress, err = NewEmailAddress(emailAddress)
		if err != nil {
			return contact, err
		}
	}
	if !util.IsEmpty(phoneNumber) {
		hasSomeContact = true
		contact.Phone.PhoneNumber, err = NewPhoneNumber(phoneNumber)
		if err != nil {
			return contact, err
		}
	}
	if !hasSomeContact {
		return contact, domain.NewInvalidEmptyParamError("email_address or phone_number")
	}

	return contact, nil
}

func (e *ContactEmail) SetUnsubscribedAt(unsubscribedAt time.Time) error {
	e.UnsubscribedAt = unsubscribedAt
	return nil
}
func (e *ContactEmail) IsUnsubscribed() bool {
	return e.UnsubscribedAt.IsZero()
}
func (e *ContactEmail) SetBouncedAt(bouncedAt time.Time) error {
	e.BouncedAt = bouncedAt
	return nil
}
func (e *ContactEmail) IsBounced() bool {
	return e.BouncedAt.IsZero()
}

func (e *ContactPhone) SetSMSUnsubscribedAt(unsubscribedAt time.Time) error {
	e.SMSUnsubscribedAt = unsubscribedAt
	return nil
}
func (e *ContactPhone) IsSMSUnsubscribed() bool {
	return e.SMSUnsubscribedAt.IsZero()
}
func (e *ContactPhone) SetWhatsAppUnsubscribedAt(bouncedAt time.Time) error {
	e.WhatsAppUnsubscribedAt = bouncedAt
	return nil
}
func (e *ContactPhone) IsWhatsAppUnsubscribed() bool {
	return e.WhatsAppUnsubscribedAt.IsZero()
}
