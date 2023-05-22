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

var AvailableBehaviorType map[BehaviorType]struct{} = map[BehaviorType]struct{}{
	BehaviorTypeSendEmail:    {},
	BehaviorTypeSendSMS:      {},
	BehaviorTypeSendWhatsapp: {},
	BehaviorTypeWaitFor:      {},
	BehaviorTypeWaitUntil:    {},
	BehaviorTypeIfAttribute:  {},
	BehaviorTypeRandom:       {},
	BehaviorTypeSplit:        {},
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

type Random struct {
}

type SendEmail struct {
}

type SendSMS struct {
}

type SendWhatsApp struct {
}

type Split struct {
}

type WaitFor struct {
}

type WaitUntil struct {
}
