package config

var (
	Columns              []string      // --columns, -C
	AllColumns           bool          // --all-columns, -A
	Limit                int32    = 0  // --limit, -L
	SortBy               string   = "" // --sort-by, -S
	ReverseSortDirection bool          // --reverse, -R
)
