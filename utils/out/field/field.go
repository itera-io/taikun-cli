package field

import (
	"fmt"
	"strings"
)

type ToStringFunc func(v interface{}) string

// Struct to represent a field in table output
type Field struct {
	// name of the field in the column
	name string

	// Corresponding JSON property name
	//
	// Putting slashes in the property name lets you access nested objects
	// Consider this JSON object:
	// {
	//   "id": 1234,
	//   "profile": {
	//     "name": "foo"
	//   }
	// }
	// Set the jsonPropertyName to "profile/name" to access the name attribute
	jsonPropertyName string

	// Function to format the field's value as a string
	toString ToStringFunc

	// Whether to visible the field in the table by default
	visible bool
}

// Create new visible Field struct
func NewVisible(name string, jsonPropertyName string) *Field {
	return &Field{
		name:             name,
		jsonPropertyName: jsonPropertyName,
		toString:         FormatByDefault,
		visible:          true,
	}
}

// Create new hidden Field struct
func NewHidden(name string, jsonPropertyName string) *Field {
	return &Field{
		name:             name,
		jsonPropertyName: jsonPropertyName,
		toString:         FormatByDefault,
		visible:          false,
	}
}

// Default field formatting function
func FormatByDefault(fieldValue interface{}) string {
	if fieldValue == nil {
		return "N/A"
	}
	if b, ok := fieldValue.(bool); ok {
		if b {
			return "Yes"
		}
		return "No"
	}
	return fmt.Sprint(fieldValue)
}

// Create new visible Field struct with To String function
func NewVisibleWithToStringFunc(name string, jsonPropertyName string, toString ToStringFunc) *Field {
	return &Field{
		name:             name,
		jsonPropertyName: jsonPropertyName,
		toString:         toString,
		visible:          true,
	}
}

// Create new hidden Field struct with To String function
func NewHiddenWithToStringFunc(name string, jsonPropertyName string, toString ToStringFunc) *Field {
	return &Field{
		name:             name,
		jsonPropertyName: jsonPropertyName,
		toString:         toString,
		visible:          false,
	}
}

// Format field value
func (f *Field) Format(value interface{}) interface{} {
	return f.toString(value)
}

// Get field's JSON property name
func (f *Field) JsonPropertyName() string {
	return f.jsonPropertyName
}

// Modifiy a field's JSON property name
func (f *Field) SetJsonPropertyName(newJsonPropertyName string) {
	f.jsonPropertyName = newJsonPropertyName
}

// Get field's name
func (f *Field) Name() string {
	return f.name
}

// Whether the field is visible
func (f *Field) IsVisible() bool {
	return f.visible
}

// Don't show field in table output
func (f *Field) Hide() {
	f.visible = false
}

// Show field in table output
func (f *Field) Show() {
	f.visible = true
}

// Whether the field's name matches the given string (case-insensitive)
func (f *Field) NameMatches(name string) bool {
	return strings.EqualFold(f.name, name)
}
