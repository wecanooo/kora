package timeutils

import "time"

// FormatTime returns a formatted datetime string (format: YYYY-MM-DD HH:MM:SS)
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FormatDate returns a formatted date string (format: YYYY-MM-DD)
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
