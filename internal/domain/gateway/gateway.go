package gateway

type GatewayManager struct {
	EmailGateway     EmailGateway
	SMSGateway       SMSGateway
	WhatsappGateway  WhatsappGateway
	SchedulerGateway SchedulerGateway
}

func NewGatewayManager(emailGateway EmailGateway, smsGateway SMSGateway, whatsappGateway WhatsappGateway, schedulerGateway SchedulerGateway) *GatewayManager {
	return &GatewayManager{
		EmailGateway:     emailGateway,
		SMSGateway:       smsGateway,
		WhatsappGateway:  whatsappGateway,
		SchedulerGateway: schedulerGateway,
	}
}
