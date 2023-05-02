package vo

import (
	"time"
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

func (e *Contact) NewContact(emailAddress string, phoneNumber string) (contact *Contact) {
	return &Contact{
		Email: ContactEmail{
			EmailAddress: EmailAddress(emailAddress),
		},
		Phone: ContactPhone{
			PhoneNumber: PhoneNumber(phoneNumber),
		},
	}
}

func (e *Contact) Validate() error {
	if err := e.Email.EmailAddress.Validate(); err != nil {
		return err
	}
	if err := e.Phone.PhoneNumber.Validate(); err != nil {
		return err
	}
	return nil
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
