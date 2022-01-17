package config

import "github.com/itera-io/taikun-cli/utils/out/fields"

var (
	Columns              []string      // --columns, -C
	Limit                int32    = 0  // --limit, -L
	SortBy               string   = "" // --sort-by, -S
	ReverseSortDirection bool          // --reverse, -R
)

func GetSortByParam(listFields fields.Fields) *string {
	param := ""
	for _, field := range listFields.AllFields() {
		if field.Matches(SortBy) {
			param = field.JsonTag()
		}
	}
	return &param
}
