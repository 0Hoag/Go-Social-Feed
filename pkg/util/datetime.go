package util

import "time"

func GetDefaultTimezone() *time.Location {
	localTimeZone, _ := time.LoadLocation("Local")
	return localTimeZone
}

func Now() time.Time {
	return time.Now().In(GetDefaultTimezone())
}
