package format

import (
	"encoding/json"
	"log"
)

const (
	prettyPrintPrefix = ""
	prettyPrintIndent = "    "
)

func prettyPrintJson(data interface{}) {
	Println(string(marshalJsonData(data)))
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
		maps[i] = jsonObjectToMap(s)
	}
	return maps
}

func jsonObjectToMap(data interface{}) map[string]interface{} {
	var m map[string]interface{}
	if err := json.Unmarshal(marshalJsonData(data), &m); err != nil {
		log.Fatal(err)
	}
	return m
}
