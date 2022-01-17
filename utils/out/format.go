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
