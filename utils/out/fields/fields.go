package fields

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/itera-io/taikun-cli/utils/out/field"
)

// Table fields
type Fields struct {
	fields []*field.Field
}

// Create new Fields struct
// Exits with error code if fields don't respect the following requirements:
// - fields must be unique within a table
// - names must be less than 20 characters
// - names must contain only uppercase letters and hyphens
// - names must not start or end with a hyphen
func New(fields []*field.Field) Fields {
	nameFrequencyMap := map[string]bool{}
	jsonTagFrequencyMap := map[string]bool{}
	for _, field := range fields {
		if !fieldNameIsValid(field.Name()) {
			log.Fatal("fields.New: Field name '", field.Name(), "' is not valid")
		}
		if nameFrequencyMap[field.Name()] {
			log.Fatal("fields.New: Field name '", field.Name(), "' is defined more than once")
		}
		nameFrequencyMap[field.Name()] = true
		if jsonTagFrequencyMap[field.JsonTag()] {
			log.Fatal("fields.New: Field JSON tag '", field.JsonTag(), "' is defined more than once")
		}
		jsonTagFrequencyMap[field.JsonTag()] = true
	}
	return Fields{
		fields: fields,
	}
}

// Returns whether or not the field's name is valid
func fieldNameIsValid(name string) bool {
	maxFieldNameLength := 20
	if len(name) == 0 || len(name) > maxFieldNameLength {
		return false
	}
	matched, err := regexp.Match("^[A-Z]+(-[A-Z]+)*$", []byte(name))
	if err != nil {
		log.Fatal("fieldNameIsValid: invalid regex pattern")
	}
	return matched
}

// Get all fields
func (f Fields) AllFields() []*field.Field {
	return f.fields
}

// Get visible fields
func (f Fields) VisibleFields() []*field.Field {
	fields := make([]*field.Field, 0)
	for _, field := range f.fields {
		if field.IsVisible() {
			fields = append(fields, field)
		}
	}
	return fields
}

// Get JSON tag of field with the given name
// Returns the tag and a boolean to indicate whether the field was found
func (f Fields) GetJsonTagFromName(name string) (jsonTag string, found bool) {
	for _, field := range f.fields {
		if field.NameMatches(name) {
			jsonTag = field.JsonTag()
			found = true
			break
		}
	}
	return
}

// Get number of visible fields
func (f Fields) VisibleSize() int {
	size := 0
	for _, field := range f.fields {
		if field.IsVisible() {
			size++
		}
	}
	return size
}

// Override the default visibility settings and display only the given fields
func (f Fields) SetVisible(fieldNames []string) {
	f.hideAll()
	for rank, fieldName := range fieldNames {
		i := f.getFieldIndex(fieldName)
		if i == -1 {
			fmt.Fprintf(os.Stderr, "Error: unknown field name '%s'\n", fieldName)
			os.Exit(1)
		} else {
			f.fields[i].Show()
			f.moveFieldBack(i, rank)
		}
	}
}

func (f Fields) getFieldIndex(fieldName string) int {
	for i, field := range f.fields {
		if field.NameMatches(fieldName) {
			return i
		}
	}
	return -1
}

func (f Fields) moveFieldBack(source int, destination int) {
	if destination > source {
		log.Fatal("Fields.moveFieldBack: destination must not be greater than source")
	}
	sourceField := f.fields[source]
	for i := source; i > destination; i-- {
		f.fields[i] = f.fields[i-1]
	}
	f.fields[destination] = sourceField
}

// Set all fields to hidden
func (f Fields) hideAll() {
	for _, field := range f.fields {
		field.Hide()
	}
}

// Set all fields to visible
func (f Fields) ShowAll() {
	for _, field := range f.fields {
		field.Show()
	}
}

// Get the list of all field names
func (f Fields) AllNames() []string {
	names := make([]string, len(f.fields))
	for i, field := range f.fields {
		names[i] = field.Name()
	}
	return names
}

// Get the list of names of the visible fields
func (f Fields) VisibleNames() []string {
	names := make([]string, 0)
	for _, field := range f.fields {
		if field.IsVisible() {
			names = append(names, field.Name())
		}
	}
	return names
}

// Get the list of JSON tags of the visible fields
func (f Fields) VisibleJsonTags() []string {
	jsonTags := make([]string, 0)
	for _, field := range f.fields {
		if field.IsVisible() {
			jsonTags = append(jsonTags, field.JsonTag())
		}
	}
	return jsonTags
}
