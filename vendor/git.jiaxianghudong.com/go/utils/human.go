package utils

import (
	"fmt"
)

// HumanTimeSecond 格式化秒数
func HumanTimeSecond(s int64, separator ...string) string {
	const (
		Minute = 60
		Hour = 60 * Minute
		Day = 24 * Hour
	)
	var d, h, m int64

	d = s / Day
	s = s % Day
	h = s / Hour
	s = s % Hour
	m = s / Minute
	s = s % Minute

	var sep string
	if len(separator) == 1 {
		sep = separator[0]
	}

	if d > 0 {
		return fmt.Sprintf("%dd%s%dh%s%dm%s%ds", d, sep, h, sep, m, sep, s)
	} else if h > 0 {
		return fmt.Sprintf("%dh%s%dm%s%ds", h, sep, m, sep, s)
	} else if m > 0 {
		return fmt.Sprintf("%dm%s%ds", m, sep, s)
	}
	return fmt.Sprintf("%ds", s)
}
