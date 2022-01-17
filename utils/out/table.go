package out

import (
	"log"
	"os"
	"reflect"

	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func printTable(data interface{}, fields fields.Fields) {
	t := newTable()

	if len(config.Columns) != 0 {
		fields.SetVisible(config.Columns)
	}

	appendHeader(t, fields.VisibleNames())

	resources := interfaceToInterfaceSlice(data)

	resourceMaps := jsonObjectsToMaps(resources)
	for _, resourceMap := range resourceMaps {
		t.AppendRow(resourceMapToRow(resourceMap, fields))
	}

	renderTable(t)
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

func resourceMapToRow(resourceMap map[string]interface{}, fields fields.Fields) []interface{} {
	row := make([]interface{}, fields.VisibleSize())
	for i, field := range fields.VisibleFields() {
		if value, found := resourceMap[field.JsonTag()]; found && value != nil {
			row[i] = field.Format(value)
		} else {
			row[i] = ""
		}
		row[i] = trimCellValue(row[i])
	}
	return row
}

func interfaceToInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		log.Fatal("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
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
