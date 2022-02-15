package out

import (
	"encoding/json"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
)

const (
	prettyPrintPrefix = ""
	prettyPrintIndent = "    "
)

func prettyPrintJson(data interface{}) error {
	jsonEncoding, err := marshalJsonData(data)
	if err != nil {
		return cmderr.ProgramError("prettyPrintJson", err)
	}
	Println(string(jsonEncoding))
	return nil
}

func getValueFromJsonMap(jsonMap map[string]interface{}, compositeKey string) (value interface{}, found bool) {
	keys := strings.Split(compositeKey, "/")
	keyCount := len(keys)
	if keyCount == 0 {
		return
	}
	value, found = jsonMap[keys[0]]
	for i := 1; i < keyCount && found; i++ {
		if m, err := jsonObjectToMap(value); err != nil {
			found = false
		} else {
			value, found = m[keys[i]]
		}
	}
	return
}

func marshalJsonData(data interface{}) ([]byte, error) {
	if data == nil {
		data = struct{}{}
	}
	return json.MarshalIndent(data, prettyPrintPrefix, prettyPrintIndent)
}

func jsonObjectsToMaps(structs []interface{}) ([]map[string]interface{}, error) {
	maps := make([]map[string]interface{}, len(structs))
	for i, s := range structs {
		m, err := jsonObjectToMap(s)
		if err != nil {
			return nil, err
		}
		maps[i] = m
	}
	return maps, nil
}

func jsonObjectToMap(data interface{}) (m map[string]interface{}, err error) {
	jsonEncoding, err := marshalJsonData(data)
	if err == nil {
		err = json.Unmarshal(jsonEncoding, &m)
	}
	return
}
