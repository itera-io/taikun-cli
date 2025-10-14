package types

import (
	"time"

	"github.com/go-openapi/strfmt"
)

const ExpectedDateFormat = "dd.mm.yyyy"
const ExpectedDateTimeFormat = "dd.mm.yyyy hh:mm:ss"

// Convert string in the format dd.mm.yyyy, dd.mm.yyyy hh:mm, or dd.mm.yyyy hh:mm:ss to a strfmt.DateTime struct
func StrToDateTime(str string) strfmt.DateTime {
	dateInRfc3339Format := strToRfc3339DateTime(str)
	myTime, _ := time.Parse(time.RFC3339, dateInRfc3339Format)

	return strfmt.DateTime(myTime)
}

// Convert string to RFC 3339 datetime format
// Supports "dd.mm.yyyy", "dd.mm.yyyy hh:mm", and "dd.mm.yyyy hh:mm:ss" formats
func strToRfc3339DateTime(date string) string {
	// Check if the string contains time (has space and colon)
	if len(date) >= 16 && date[10] == ' ' && date[13] == ':' {
		// Check if it has seconds (format: dd.mm.yyyy hh:mm:ss)
		if len(date) >= 19 && date[16] == ':' {
			hour := date[11:13]
			minute := date[14:16]
			second := date[17:19]
			return date[6:10] + "-" + date[3:5] + "-" + date[0:2] + "T" + hour + ":" + minute + ":" + second + "Z"
		} else {
			// Format: dd.mm.yyyy hh:mm (default seconds to 00)
			hour := date[11:13]
			minute := date[14:16]
			return date[6:10] + "-" + date[3:5] + "-" + date[0:2] + "T" + hour + ":" + minute + ":00Z"
		}
	} else {
		// Format: dd.mm.yyyy (default to midnight)
		return date[6:10] + "-" + date[3:5] + "-" + date[0:2] + "T00:00:00Z"
	}
}

// Whether `str` is a valid date in the format dd.mm.yyyy, dd.mm.yyyy hh:mm, or dd.mm.yyyy hh:mm:ss
func StrIsValidDate(str string) bool {
	// Check for date-only format (dd.mm.yyyy)
	if len(str) == 10 {
		dateInRfc3339Format := strToRfc3339DateTime(str)
		_, err := time.Parse(time.RFC3339, dateInRfc3339Format)
		return err == nil
	}

	// Check for datetime format with minutes (dd.mm.yyyy hh:mm)
	if len(str) == 16 && str[10] == ' ' && str[13] == ':' {
		dateInRfc3339Format := strToRfc3339DateTime(str)
		_, err := time.Parse(time.RFC3339, dateInRfc3339Format)
		return err == nil
	}

	// Check for datetime format with seconds (dd.mm.yyyy hh:mm:ss)
	if len(str) == 19 && str[10] == ' ' && str[13] == ':' && str[16] == ':' {
		dateInRfc3339Format := strToRfc3339DateTime(str)
		_, err := time.Parse(time.RFC3339, dateInRfc3339Format)
		return err == nil
	}

	return false
}
