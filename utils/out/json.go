package out

import (
	"encoding/json"
	"log"
	"strings"
)

const (
	prettyPrintPrefix = ""
	prettyPrintIndent = "    "
)

func prettyPrintJson(data interface{}) {
	Println(string(marshalJsonData(data)))
}

func getValueFromJsonMap(m map[string]interface{}, compositeKey string) (value interface{}, found bool) {
	keys := strings.Split(compositeKey, "/")
	keyCount := len(keys)
	if keyCount == 0 {
		return
	}
	value, found = m[keys[0]]
	for i := 1; i < keyCount && found; i++ {
		if m, err := jsonObjectToMap(value); err != nil {
			found = false
		} else {
			value, found = m[keys[i]]
		}
	}
	return
}

func marshalJsonData(data interface{}) []byte {
	if data == nil {
		data = struct{}{}
	}
	jsonBytes, err := json.MarshalIndent(data, prettyPrintPrefix, prettyPrintIndent)
	if err != nil {
		log.Fatal(err)
	}
	return jsonBytes
}

func jsonObjectsToMaps(structs []interface{}) []map[string]interface{} {
	maps := make([]map[string]interface{}, len(structs))
	for i, s := range structs {
		m, err := jsonObjectToMap(s)
		if err != nil {
			log.Fatal(err)
		}
		maps[i] = m
	}
	return maps
}

func jsonObjectToMap(data interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(marshalJsonData(data), &m); err != nil {
		return nil, err
	}
	return m, nil
}
