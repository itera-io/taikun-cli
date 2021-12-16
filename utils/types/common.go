package types

import "fmt"

func MapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func MapContains(m map[string]interface{}, key string) bool {
	_, contains := m[key]
	return contains
}

func UnknownFlagValueError(flag string, received string, expected []string) error {
	return fmt.Errorf("Unknown %s: %s, expected one of %v.", flag, received, expected)
}
