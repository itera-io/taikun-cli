package options

const (
	OutputFormatJson  = "json"
	OutputFormatTable = "table"
)

type Outputter interface {
	GetOutputFormatOption() *string
	GetQuietOption() *bool
}

func OutputFormatIsValid(outputter Outputter) bool {
	switch *outputter.GetOutputFormatOption() {
	case OutputFormatJson:
		return true
	case OutputFormatTable:
		return true
	default:
		return false
	}
}
