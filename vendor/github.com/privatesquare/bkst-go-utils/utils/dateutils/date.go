package dateutils

import "time"

const (
	dateTimeFormat = "2006-01-02 15:04:05"
)

// GetDateTimeNow returns the current data time in UTC format
func GetDateTimeNow() time.Time {
	return time.Now().UTC()
}

// GetDateTimeNowFormat returns the current time in the "2006-01-02 15:04:05" format
func GetDateTimeNowFormat() string {
	return GetDateTimeNow().Format(dateTimeFormat)
}
