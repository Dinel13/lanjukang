package utilities

import "time"

// ConvertToDate converts a string to a date.
func ConvertToDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

// ConvertToDateTime converts a string to a date time.
func ConvertToDateTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", date)
}

// ConvertToTIMEHOUR converts a string to a time hour.
func ConvertToTimeHour(date string) (time.Time, error) {
	return time.Parse("15:04", date)
}
