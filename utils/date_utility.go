package utils

import (
	"regexp"
	"time"
)

var hexPattern = regexp.MustCompile("^[0-9a-fA-F]+$")

// IsHexString checks if a string is a valid hexadecimal representation.
func IsHexString(str string) bool {
	return hexPattern.MatchString(str)
}

// DateWithYearMonthDay returns the current date in the "yyyyMMdd" format.
func DateWithYearMonthDay() string {
	date := time.Now()
	return date.Format("20060102")
}
