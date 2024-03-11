package common

import (
	"fmt"
	"time"
)

func FormatDateTime(t time.Time) string {
	year, month, day := t.Date()
	hour, minute, _ := t.Clock()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d", year, month, day, hour, minute)
}

func FormatDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
