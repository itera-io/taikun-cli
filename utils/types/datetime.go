package types

import (
	"time"

	"github.com/go-openapi/strfmt"
)

const ExpectedDateFormat = "dd.mm.yyyy"

// Convert string in the format dd.mm.yyyy to a strfmt.DateTime struct
func StrToDateTime(str string) strfmt.DateTime {
	dateInRfc3339Format := strToRfc3339DateTime(str)
	myTime, _ := time.Parse(time.RFC3339, dateInRfc3339Format)

	return strfmt.DateTime(myTime)
}

// Convert string to RFC 3339 datetime format
func strToRfc3339DateTime(date string) string {
	return date[6:10] + "-" + date[3:5] + "-" + date[0:2] + "T00:00:00Z"
}

// Whether `str` is a valid date in the format dd.mm.yyyy
func StrIsValidDate(str string) bool {
	if len(str) != 10 {
		return false
	}

	dateInRfc3339Format := strToRfc3339DateTime(str)
	_, err := time.Parse(time.RFC3339, dateInRfc3339Format)

	return err == nil
}
