package fields

import (
	"fmt"
	"log"
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

// Get JSON tag of field with the given name
// Returns empty string of no field has the given name
func (f Fields) GetJsonTagFromName(name string) string {
	for _, field := range f.fields {
		if field.NameMatches(name) {
			return field.JsonTag()
		}
	}
	return ""
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
