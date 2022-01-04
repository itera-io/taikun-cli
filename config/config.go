package config

func init() {
	OutputFormat = OutputFormatTable
}

// root command's persistent flags
var (
	NoDecorate      bool   // --no-decorate
	OutputFormat    string // --format, -F
	Quiet           bool   // --quiet, -q
	ShowLargeValues bool   // --show-large-values
)

const (
	OutputFormatJson  = "json"
	OutputFormatTable = "table"
)

var OutputFormats = []string{
	OutputFormatJson,
	OutputFormatTable,
}

func OutputFormatIsValid() bool {
	return OutputFormat == OutputFormatJson ||
		OutputFormat == OutputFormatTable
}
