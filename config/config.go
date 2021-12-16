package config

func init() {
	OutputFormat = OutputFormatTable
}

var OutputFormat string

var ShowLargeValues bool

const (
	OutputFormatJson  = "json"
	OutputFormatTable = "table"
)

func OutputFormatIsValid() bool {
	return OutputFormat == OutputFormatJson ||
		OutputFormat == OutputFormatTable
}
