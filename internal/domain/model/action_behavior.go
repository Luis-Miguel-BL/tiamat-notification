package model

type BehaviorType string

const (
	BehaviorTypeSendEmail    BehaviorType = "send-email"
	BehaviorTypeSendSMS      BehaviorType = "send-sms"
	BehaviorTypeSendWhatsapp BehaviorType = "send-whatsapp"
	BehaviorTypeWaitFor      BehaviorType = "wait-for"
	BehaviorTypeWaitUntil    BehaviorType = "wait-until"
	BehaviorTypeIfAttribute  BehaviorType = "if-attribute"
	BehaviorTypeRandom       BehaviorType = "random"
	BehaviorTypeSplit        BehaviorType = "split"
)

type ValidateBehaviorFunc func(ActionBehavior) bool

var AvailableBehaviorType = map[BehaviorType]ValidateBehaviorFunc{
	BehaviorTypeSendEmail:    validateSendEmail,
	BehaviorTypeSendSMS:      validateSendSMS,
	BehaviorTypeSendWhatsapp: validateSendWhatsapp,
	BehaviorTypeWaitFor:      validateWaitFor,
	BehaviorTypeWaitUntil:    validateWaitUntil,
	BehaviorTypeIfAttribute:  validateIfAttribute,
	BehaviorTypeRandom:       validateRandom,
	BehaviorTypeSplit:        validateSplit,
}

type ActionBehavior struct {
	Type         BehaviorType
	SendEmail    SendEmail
	SendSMS      SendSMS
	SendWhatsApp SendWhatsApp
	WaitFor      WaitFor
	WaitUntil    WaitUntil
	IfAttribute  IfAttribute
	Random       Random
	Split        Split
}

type IfAttribute struct {
}

func validateIfAttribute(behavior ActionBehavior) bool {
	return true
}

type Random struct {
}

func validateRandom(behavior ActionBehavior) bool {
	return true
}

type SendEmail struct {
}

func validateSendEmail(behavior ActionBehavior) bool {
	return true
}

type SendSMS struct {
}

func validateSendSMS(behavior ActionBehavior) bool {
	return true
}

type SendWhatsApp struct {
}

func validateSendWhatsapp(behavior ActionBehavior) bool {
	return true
}

type Split struct {
}

func validateSplit(behavior ActionBehavior) bool {
	return true
}

type WaitFor struct {
}

func validateWaitFor(behavior ActionBehavior) bool {
	return true
}

type WaitUntil struct {
}

func validateWaitUntil(behavior ActionBehavior) bool {
	return true
}
