package out

import (
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmdutils/options"
)

const maxColumnWidth = 50
const trimmedValueSuffix = "..."

func trimCellValue(opts options.TableWriter, value interface{}) interface{} {
	if !*opts.GetShowLargeValuesOption() {
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

func FormatDateTimeString(v interface{}) string {
	dateTime := v.(string)
	dateTime = strings.Replace(dateTime, "T", " ", 1)
	dateTime = strings.Replace(dateTime, "Z", "", 1)
	return dateTime
}

func FormatLockStatus(v interface{}) string {
	lockStatus := v.(bool)
	if lockStatus {
		return "Locked"
	}
	return "Unlocked"
}
