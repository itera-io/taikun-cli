package fields

import (
	"fmt"
	"os"

	"github.com/itera-io/taikun-cli/utils/out/field"
)

// Table fields
type Fields struct {
	fields []*field.Field
}

// Create new Fields struct
func New(fields []*field.Field) Fields {
	return Fields{
		fields: fields,
	}
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
	for _, fieldName := range fieldNames {
		field := f.get(fieldName)
		if field == nil {
			fmt.Fprintf(os.Stderr, "Error: unknown field name '%s'\n", fieldName)
			os.Exit(1)
		} else {
			field.Show()
		}
	}
}

func (f Fields) hideAll() {
	for _, field := range f.fields {
		field.Hide()
	}
}

func (f Fields) ShowAll() {
	for _, field := range f.fields {
		field.Show()
	}
}

func (f Fields) get(fieldName string) *field.Field {
	for _, field := range f.fields {
		if field.Matches(fieldName) {
			return field
		}
	}
	return nil
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
