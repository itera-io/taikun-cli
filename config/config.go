package config

import "errors"

func init() {
	OutputFormat = OutputFormatTable
}

var OutputFormat string
var OutputFormatInvalidError = errors.New("Unknown output format")

var ShowLargeValues bool

const (
	OutputFormatJson  = "json"
	OutputFormatTable = "table"
)

func OutputFormatIsValid() bool {
	return OutputFormat == OutputFormatJson ||
		OutputFormat == OutputFormatTable
}
