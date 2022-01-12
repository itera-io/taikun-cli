package out

import (
	"github.com/itera-io/taikun-cli/config"
)

func PrintResults(slice interface{}, fields ...string) {
	if config.OutputFormat == config.OutputFormatJson {
		prettyPrintJson(slice)
	} else if config.OutputFormat == config.OutputFormatTable {
		prettyPrintTable(slice, fields...)
	}
}
