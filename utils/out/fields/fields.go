package fields

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out/field"
)

// Table fields
type Fields struct {
	fields           []*field.Field
	parentObjectName string
}

// Create new Fields struct
// Exits with error code if fields don't respect the following requirements:
// - fields must be unique within a table
// - names must be less than 20 characters
// - names must contain only uppercase letters and hyphens
// - names must not start or end with a hyphen
func New(fields []*field.Field) Fields {
	nameFrequencyMap := map[string]bool{}
	jsonPropertyNameFreqMap := map[string]bool{}
	for _, field := range fields {
		if !fieldNameIsValid(field.Name()) {
			panic(fmt.Sprintf("fields.New: Field name '%s' is not valid", field.Name()))
		}
		if nameFrequencyMap[field.Name()] {
			panic(fmt.Sprintf("fields.New: Field name '%s' is defined more than once", field.Name()))
		}
		nameFrequencyMap[field.Name()] = true
		if jsonPropertyNameFreqMap[field.JsonPropertyName()] {
			panic(fmt.Sprintf("fields.New: Field JSON property name '%s' is defined more than once", field.JsonPropertyName()))
		}
		jsonPropertyNameFreqMap[field.JsonPropertyName()] = true
	}
	return Fields{
		fields: fields,
	}
}

// Same as New, create new Fields struct but fields are in a nested object
func NewNested(fields []*field.Field, parentObjectName string) Fields {
	f := New(fields)
	f.parentObjectName = parentObjectName
	return f
}

// Returns whether the fields belong to a nested structure and the parent
// object's name
func (f Fields) AreNested() (parentObjectName string, areNested bool) {
	return f.parentObjectName, f.parentObjectName != ""
}

// Modify the JSON property name of the field with the given name
// If no field is found with the given name, returns an error
func (f Fields) SetFieldJsonPropertyName(name string, jsonPropertyName string) error {
	for _, field := range f.fields {
		if field.NameMatches(name) {
			field.SetJsonPropertyName(jsonPropertyName)
			return nil
		}
	}
	return cmderr.ProgramError("SetFieldJsonPropertyName", fmt.Errorf("unknown field name: %s", name))
}

// Returns whether or not the field's name is valid
func fieldNameIsValid(name string) bool {
	maxFieldNameLength := config.MaxCellWidth
	if len(name) == 0 || len(name) > maxFieldNameLength {
		return false
	}
	matched, err := regexp.Match("^[A-Z0-9]+(-[A-Z0-9]+)*$", []byte(name))
	if err != nil {
		panic("fieldNameIsValid: invalid regex pattern")
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

// Get JSON property name of field with the given name
// Returns the property name and a boolean to indicate whether the field was found
func (f Fields) GetJsonPropertyNameFromName(name string) (jsonPropertyName string, found bool) {
	for _, field := range f.fields {
		if field.NameMatches(name) {
			jsonPropertyName = field.JsonPropertyName()
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
func (f Fields) SetVisible(fieldNames []string) error {
	f.hideAll()
	for rank, fieldName := range fieldNames {
		i := f.getFieldIndex(fieldName)
		if i == -1 {
			return fmt.Errorf("Error: unknown field name '%s'\n", fieldName)
		} else {
			f.fields[i].Show()
			if err := f.moveFieldBack(i, rank); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f Fields) getFieldIndex(fieldName string) int {
	for i, field := range f.fields {
		if field.NameMatches(fieldName) {
			return i
		}
	}
	return -1
}

func (f Fields) moveFieldBack(source int, destination int) error {
	if destination > source {
		return errors.New("Fields.moveFieldBack: destination must not be greater than source")
	}
	sourceField := f.fields[source]
	for i := source; i > destination; i-- {
		f.fields[i] = f.fields[i-1]
	}
	f.fields[destination] = sourceField
	return nil
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

// Get the list of JSON property names of the visible fields
func (f Fields) VisibleFieldsJsonPropertyNames() []string {
	jsonPropertyNames := make([]string, 0)
	for _, field := range f.fields {
		if field.IsVisible() {
			jsonPropertyNames = append(jsonPropertyNames, field.JsonPropertyName())
		}
	}
	return jsonPropertyNames
}
