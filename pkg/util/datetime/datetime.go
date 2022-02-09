package datetime

import "time"

// FormatDate returns time formatted as RFC3339 (2021-11-10T17:41:48+09:00).
func FormatDate(t time.Time) string {
	return t.Format(time.RFC3339)
}

// AddDate adds date.
func AddDate(date int) time.Time {
	t := time.Now()

	return t.AddDate(0, 0, date)
}

// StartOfDay returns the start of a day for the given date. The result will be in the local timezone.
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// EndOfDay returns the end of a day for the given date. The result will be in the local timezone.
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local)
}
