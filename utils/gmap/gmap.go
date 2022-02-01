package gmap

import (
	"fmt"
	"strings"
)

type GenericMap struct {
	m map[string]interface{}
}

func New(m map[string]interface{}) GenericMap {
	for key := range m {
		if key != strings.ToLower(key) {
			panic(fmt.Sprintf("GenericMap keys must be lowercase, have: %s", key))
		}
	}
	return GenericMap{
		m: m,
	}
}

// Search for value associated to the given key, is case-insensitive
func (m GenericMap) Get(key string) interface{} {
	return m.m[strings.ToLower(key)]
}

func (m GenericMap) Keys() []string {
	keys := make([]string, 0, len(m.m))
	for key := range m.m {
		keys = append(keys, key)
	}
	return keys
}

// Check whether key exists, is case-insensitive
func (m GenericMap) Contains(key string) bool {
	_, contains := m.m[strings.ToLower(key)]
	return contains
}
