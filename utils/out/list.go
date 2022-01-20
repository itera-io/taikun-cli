package out

import (
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/fields"
)

func PrintResults(slice interface{}, fields fields.Fields) {
	if config.OutputFormat == config.OutputFormatJson {
		prettyPrintJson(slice)
	} else if config.OutputFormat == config.OutputFormatTable {
		printTable(slice, fields)
	}
}
