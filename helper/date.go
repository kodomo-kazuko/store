package helper

import (
	"strings"
	"time"
)

const (
	mysqlDateTimeFormat = "2006-01-02 15:04:05"
	mysqlDateFormat     = "2006-01-02"
)

func ToDate(value string, endOfDay bool) (time.Time, error) {
	// Remove % from the string
	value = strings.Replace(value, "%", "", -1)

	// Check if the input string contains time values
	if len(value) == len(mysqlDateFormat) {
		if endOfDay {
			// If endOfDay is true, add 23:59:00
			value = value + " 23:59:59"
		} else {
			// If the time values are not provided, add 00:00:00
			value = value + " 00:00:00"
		}
	}

	// Parse the date
	date, err := time.Parse(mysqlDateTimeFormat, value)
	if err != nil {
		return time.Time{}, err // Return zero value of time.Time on error
	}

	return date, nil
}
