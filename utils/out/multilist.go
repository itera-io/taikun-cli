package out

import (
	"log"

	"github.com/itera-io/taikun-cli/cmd/cmdutils/options"
	"github.com/itera-io/taikun-cli/utils/out/fields"
)

// Allows printing of resources of different types into one table.
// If resource types is not empty,
// a 'type' column will be added to the table,
// the value of the 'type' cell for the resources
// contained in the slice at index *i* of *resourceSlices*
// will be the type at index *i* of *resourceTypes*.
// Thus, *resourceSlices* and *resourceTypes* MUST have the same length.
func PrintResultsOfDifferentTypes(
	resourceSlices []interface{},
	resourceTypes []string,
	opts interface{},
	fields fields.Fields,
) {
	outputOpts := opts.(options.Outputter)
	if *outputOpts.GetOutputFormatOption() == options.OutputFormatJson {
		for _, slice := range resourceSlices {
			prettyPrintJson(outputOpts, slice)
		}
	} else if *outputOpts.GetOutputFormatOption() == options.OutputFormatTable {
		printTableWithDifferentTypes(opts.(options.TableWriter), resourceSlices, resourceTypes, fields)
	}
}

func printTableWithDifferentTypes(
	opts interface{},
	resourceSlices []interface{},
	resourceTypes []string,
	fields fields.Fields,
) {
	addTypeColumn := true
	if len(resourceTypes) == 0 {
		addTypeColumn = false
	} else if len(resourceSlices) != len(resourceTypes) {
		log.Fatal("PrintMultipleResults: resourcesSlices and resourceTypes must have the same length")
	}

	tableOpts := opts.(options.TableWriter)
	t := newTable(tableOpts)

	if *tableOpts.GetAllColumnsOption() {
		fields.ShowAll()
		addTypeColumn = true
	} else if len(*tableOpts.GetColumnsOption()) != 0 {
		fields.SetVisible(*tableOpts.GetColumnsOption())
		addTypeColumn = false
	}

	header := fields.VisibleNames()
	if addTypeColumn {
		header = append(header, "TYPE")
	}
	appendHeader(tableOpts, t, header)

	for resourceIndex, resourcesData := range resourceSlices {
		resources := resourcesData.([]interface{})
		resourceMaps := jsonObjectsToMaps(resources)
		for _, resourceMap := range resourceMaps {
			row := resourceMapToRow(tableOpts, resourceMap, fields)
			if addTypeColumn {
				row = append(row, resourceTypes[resourceIndex])
			}
			t.AppendRow(row)
		}
	}

	renderTable(opts.(options.Outputter), t)
}
