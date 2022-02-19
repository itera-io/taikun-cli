package out

import (
	"errors"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/config"
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
	fields fields.Fields,
) error {
	if config.OutputFormat == config.OutputFormatJson {
		for _, slice := range resourceSlices {
			if err := prettyPrintJson(slice); err != nil {
				return err
			}
		}

		return nil
	}

	return printTableWithDifferentTypes(resourceSlices, resourceTypes, fields)
}

func printTableWithDifferentTypes(
	resourceSlices []interface{},
	resourceTypes []string,
	fields fields.Fields,
) error {
	addTypeColumn := true
	if len(resourceTypes) == 0 {
		addTypeColumn = false
	} else if len(resourceSlices) != len(resourceTypes) {
		return cmderr.ProgramError("PrintMultipleResults", errors.New("resourcesSlices and resourceTypes must have the same length"))
	}

	tab := newTable()

	if config.AllColumns {
		addTypeColumn = true

		fields.ShowAll()
	} else if len(config.Columns) != 0 {
		if err := fields.SetVisible(config.Columns); err != nil {
			return err
		}
		addTypeColumn = false
	}

	header := fields.VisibleNames()
	if addTypeColumn {
		header = append(header, "TYPE")
	}

	appendHeader(tab, header)

	for resourceIndex, resourcesData := range resourceSlices {
		if resourceIndex > 0 {
			appendSeparator(tab)
		}

		resources := resourcesData.([]interface{})

		resourceMaps, err := jsonObjectsToMaps(resources)
		if err != nil {
			return cmderr.ProgramError("printTableWithDifferentTypes", err)
		}

		for _, resourceMap := range resourceMaps {
			row := resourceMapToRow(resourceMap, fields)
			if addTypeColumn {
				row = append(row, resourceTypes[resourceIndex])
			}

			tab.AppendRow(row)
		}
	}

	renderTable(tab)

	return nil
}
