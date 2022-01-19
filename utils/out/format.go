package out

import (
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/config"
)

const maxColumnWidth = 50
const trimmedValueSuffix = "..."

func trimCellValue(value interface{}) interface{} {
	if !config.ShowLargeValues {
		if str, isString := value.(string); isString {
			if len(str) > maxColumnWidth {
				str = str[:(maxColumnWidth - len(trimmedValueSuffix))]
				str += trimmedValueSuffix
			}
			return str
		}
	}
	return value
}

func resourceIDToString(id interface{}) string {
	if str, isString := id.(string); isString {
		return strings.ReplaceAll(str, "\"", "")
	}
	return fmt.Sprint(id)
}

// Format a datetime string with format '<YYYY>-<MM>-<DD>T<HH>:<MM>:<SS>Z'
func FormatDateTimeString(v interface{}) string {
	if dateTime, ok := v.(string); ok {
		if dateTime == "" {
			return "N/A"
		}
		dateTime = strings.Replace(dateTime, "T", " ", 1)
		dateTime = strings.Replace(dateTime, "Z", "", 1)
		return dateTime
	}
	return "N/A"

}

// Display true/false as Locked/Unlocked
func FormatLockStatus(v interface{}) string {
	if lockStatus, ok := v.(bool); ok {
		if lockStatus {
			return "Locked"
		}
		return "Unlocked"
	}
	return "N/A"
}

// If not availabale, display N/A
func FormatProjectHealth(v interface{}) string {
	if health, ok := v.(string); ok {
		if health == "None" {
			return "N/A"
		}
		return health
	}
	return "N/A"
}
