package gateway

import "time"

type SchedulerGateway interface {
	Scheduler(timestamp time.Time)
}
