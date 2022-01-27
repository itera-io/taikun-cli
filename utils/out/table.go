package out

import (
	"errors"
	"os"
	"reflect"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func printTable(data interface{}, fields fields.Fields) error {
	t := newTable()

	if config.AllColumns {
		fields.ShowAll()
	} else if len(config.Columns) != 0 {
		if err := fields.SetVisible(config.Columns); err != nil {
			return cmderr.ProgramError("printTable", err)
		}
	}

	appendHeader(t, fields.VisibleNames())

	resources, err := interfaceToInterfaceSlice(data)
	if err != nil {
		return cmderr.ProgramError("printTable", err)
	}

	if parentObjectName, nested := fields.AreNested(); nested {
		allNestedResources := make([]interface{}, 0)
		for i := range resources {
			nestedResources, err := getNestedResources(resources[i], parentObjectName)
			if err != nil {
				return err
			}
			allNestedResources = append(allNestedResources, nestedResources...)
		}
		resources = allNestedResources
	}

	resourceMaps, err := jsonObjectsToMaps(resources)
	if err != nil {
		return cmderr.ProgramError("printTable", err)
	}
	for _, resourceMap := range resourceMaps {
		t.AppendRow(resourceMapToRow(resourceMap, fields))
	}

	renderTable(t)
	return nil
}

func newTable() table.Writer {
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleDefault)
	t.Style().Format.Header = text.FormatDefault
	t.Style().Options = table.OptionsNoBorders

	if config.NoDecorate {
		t.Style().Options = table.OptionsNoBordersAndSeparators
		t.Style().Box.PaddingLeft = ""
	}

	return t
}

func getNestedResources(resource interface{}, parentObjectName string) (nestedResources []interface{}, err error) {
	resourceMap, err := jsonObjectToMap(resource)
	if err != nil {
		return nil, cmderr.ProgramError("getNestedResource", err)
	}
	nestedData, ok := resourceMap[parentObjectName]
	if !ok {
		return nil, cmderr.ProgramError("getNestedResource", errors.New("could not find nested resource"))
	}
	nestedResources, ok = nestedData.([]interface{})
	if !ok {
		return nil, cmderr.ProgramError("getNestedResource", errors.New("nested resource is not array"))
	}
	return
}

func resourceMapToRow(resourceMap map[string]interface{}, fields fields.Fields) []interface{} {
	row := make([]interface{}, fields.VisibleSize())
	for i, field := range fields.VisibleFields() {
		if value, found := getValueFromJsonMap(resourceMap, field.JsonPropertyName()); found && value != nil {
			row[i] = field.Format(value)
		} else {
			row[i] = ""
		}
		row[i] = trimCellValue(row[i])
	}
	return row
}

func interfaceToInterfaceSlice(slice interface{}) ([]interface{}, error) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("failed to convert interface to interface slice")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil, nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret, nil
}

func stringSliceToRow(fields []string) table.Row {
	row := make([]interface{}, len(fields))
	for i, field := range fields {
		row[i] = field
	}
	return row
}

func appendHeader(t table.Writer, fields []string) {
	if !config.NoDecorate {
		t.AppendHeader(stringSliceToRow(fields))
	}
}

func appendSeparator(t table.Writer) {
	if !config.NoDecorate {
		t.AppendSeparator()
	}
}
