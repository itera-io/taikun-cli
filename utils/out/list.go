package out

import (
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/fields"
)

func PrintResults(slice interface{}, fields fields.Fields) error {
	if config.OutputFormat == config.OutputFormatJson {
		return prettyPrintJson(slice)
	}

	return printTable(slice, fields)
}
