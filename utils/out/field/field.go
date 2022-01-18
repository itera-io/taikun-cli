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

	// Corresponding JSON tag
	jsonTag string

	// Function to format the field's value as a string
	toString ToStringFunc

	// Whether to visible the field in the table by default
	visible bool
}

// Create new visible Field struct
func NewVisible(name string, jsonTag string) *Field {
	return &Field{
		name:     name,
		jsonTag:  jsonTag,
		toString: FormatByDefault,
		visible:  true,
	}
}

// Create new hidden Field struct
func NewHidden(name string, jsonTag string) *Field {
	return &Field{
		name:     name,
		jsonTag:  jsonTag,
		toString: FormatByDefault,
		visible:  false,
	}
}

func FormatByDefault(v interface{}) string {
	if v == nil {
		return ""
	}
	if b, ok := v.(bool); ok {
		if b {
			return "Yes"
		}
		return "No"
	}
	return fmt.Sprint(v)
}

// Create new visible Field struct with To String function
func NewVisibleWithToStringFunc(name string, jsonTag string, toString ToStringFunc) *Field {
	return &Field{
		name:     name,
		jsonTag:  jsonTag,
		toString: toString,
		visible:  true,
	}
}

// Create new hidden Field struct with To String function
func NewHiddenWithToStringFunc(name string, jsonTag string, toString ToStringFunc) *Field {
	return &Field{
		name:     name,
		jsonTag:  jsonTag,
		toString: toString,
		visible:  false,
	}
}

// Format field value
func (f *Field) Format(value interface{}) interface{} {
	return f.toString(value)
}

// Get field's JSON tag
func (f *Field) JsonTag() string {
	return f.jsonTag
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

// Whether the field matches the given identifying string
func (f *Field) Matches(id string) bool {
	id = strings.ToLower(id)
	return strings.ToLower(f.name) == id || strings.ToLower(f.jsonTag) == id
}
