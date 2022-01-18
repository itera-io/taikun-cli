package out

import (
	"fmt"
	"os"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils/options"
	"github.com/itera-io/taikun-cli/utils/out/fields"
)

func PrintResult(resource interface{}, opts interface{}, fields fields.Fields) {
	createOpts := opts.(options.Creator)
	if *createOpts.GetOutputOnlyIDOption() {
		printResourceID(opts.(options.Outputter), resource)
	} else {
		outputOpts := opts.(options.Outputter)
		if *outputOpts.GetOutputFormatOption() == options.OutputFormatJson {
			prettyPrintJson(outputOpts, resource)
		} else if *outputOpts.GetOutputFormatOption() == options.OutputFormatTable {
			printApiResponseTable(opts.(options.TableWriter), resource, fields)
		}
	}
}

func printResourceID(opts options.Outputter, resource interface{}) {
	resourceMap := jsonObjectToMap(resource)
	if id, found := resourceMap["id"]; found {
		Println(opts, resourceIDToString(id))
	} else {
		fmt.Fprintln(os.Stderr, "ID not found")
	}
}

func printApiResponseTable(opts interface{}, response interface{}, fields fields.Fields) {
	tableOpts := opts.(options.TableWriter)
	if *tableOpts.GetAllColumnsOption() {
		fields.ShowAll()
	} else if len(*tableOpts.GetColumnsOption()) != 0 {
		fields.SetVisible(*tableOpts.GetColumnsOption())
	}

	resourceMap := getApiResponseResourceMap(response)

	t := newTable(tableOpts)
	for _, field := range fields.VisibleFields() {
		t.AppendRow([]interface{}{
			field.Name(),
			field.Format(resourceMap[field.JsonTag()]),
		})
	}

	renderTable(opts.(options.Outputter), t)
}

func getApiResponseResourceMap(response interface{}) map[string]interface{} {
	resourceMap := jsonObjectToMap(response)
	if resourceMap[api.ResultField] != nil {
		resourceMap = resourceMap[api.ResultField].(map[string]interface{})
	} else if resourceMap[api.PayloadField] != nil {
		resourceMap = resourceMap[api.PayloadField].(map[string]interface{})
	}
	return resourceMap
}
