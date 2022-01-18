package out

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils/options"
	"github.com/itera-io/taikun-cli/utils/out/fields"
)

func PrintResults(slice interface{}, opts interface{}, fields fields.Fields) {
	outputOpts := opts.(options.Outputter)
	if *outputOpts.GetOutputFormatOption() == options.OutputFormatJson {
		prettyPrintJson(outputOpts, slice)
	} else if *outputOpts.GetOutputFormatOption() == options.OutputFormatTable {
		printTable(opts.(options.TableWriter), slice, fields)
	}
}
