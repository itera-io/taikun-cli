package out

import (
	"fmt"
	"os"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/fields"
)

func PrintResult(resource interface{}, fields fields.Fields) {
	if config.OutputOnlyID {
		printResourceID(resource)
	} else {
		if config.OutputFormat == config.OutputFormatJson {
			prettyPrintJson(resource)
		} else if config.OutputFormat == config.OutputFormatTable {
			printApiResponseTable(resource, fields)
		}
	}
}

func printResourceID(resource interface{}) {
	resourceMap := jsonObjectToMap(resource)
	if id, found := resourceMap["id"]; found {
		Println(resourceIDToString(id))
	} else {
		fmt.Fprintln(os.Stderr, "ID not found")
	}
}

func printApiResponseTable(response interface{}, fields fields.Fields) {
	if config.AllColumns {
		fields.ShowAll()
	} else if len(config.Columns) != 0 {
		fields.SetVisible(config.Columns)
	}

	resourceMap := getApiResponseResourceMap(response)

	t := newTable()
	for _, field := range fields.VisibleFields() {
		t.AppendRow([]interface{}{
			field.Name(),
			field.Format(resourceMap[field.JsonTag()]),
		})
	}

	renderTable(t)
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
