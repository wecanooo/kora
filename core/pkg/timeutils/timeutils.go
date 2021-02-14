package timeutils

import "time"

// FormatTime YYYY-MM-DD HH:MM:SS 형태 반환
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FormatDate YYYY-MM-DD 형태 반환
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
