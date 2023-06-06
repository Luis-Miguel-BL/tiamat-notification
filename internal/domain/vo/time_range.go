package vo

import (
	"time"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/domain"
)

type TimeRange struct {
	availableTime map[time.Weekday][]Hour
}

type Hour int

const (
	Hour0 Hour = iota
	Hour1
	Hour2
	Hour3
	Hour4
	Hour5
	Hour6
	Hour7
	Hour8
	Hour9
	Hour10
	Hour11
	Hour12
	Hour13
	Hour14
	Hour15
	Hour16
	Hour17
	Hour18
	Hour19
	Hour20
	Hour21
	Hour22
	Hour23
)

func NewTimeRange(availableTime map[time.Weekday][]int) (timeRange TimeRange, err error) {
	for weekDay, availableHours := range availableTime {
		timeRange.availableTime[weekDay] = []Hour{}
		for _, availableHour := range availableHours {
			if availableHour < 0 || availableHour > 23 {
				continue
			}
			timeRange.availableTime[weekDay] = append(timeRange.availableTime[weekDay], Hour(availableHour))
		}
		if len(timeRange.availableTime[weekDay]) == 0 {
			delete(timeRange.availableTime, weekDay)
		}
	}
	if len(timeRange.availableTime) == 0 {
		return timeRange, domain.NewInvalidEmptyParamError("time-range")
	}
	return timeRange, nil
}

func (vo TimeRange) IsAvailable(timestamp time.Time) bool {
	availableHours, ok := vo.availableTime[timestamp.Weekday()]
	if ok {
		for _, avavailableHour := range availableHours {
			if avavailableHour == Hour(timestamp.Hour()) {
				return true
			}
		}
	}

	return false
}

func (vo TimeRange) NextAvailableTime(currentTime time.Time) (timestamp time.Time) {
	if vo.IsAvailable(currentTime) {
		return currentTime
	}

	availableHours, ok := vo.availableTime[currentTime.Weekday()]
	if ok {
		for _, availableHour := range availableHours {
			if int(availableHour) > currentTime.Hour() {
				return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), int(availableHour), 0, 0, 0, time.UTC)
			}
		}
	}

	nextDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Add(time.Hour*24).Day(), 0, 0, 0, 0, time.UTC)
	return vo.NextAvailableTime(nextDay)

}
