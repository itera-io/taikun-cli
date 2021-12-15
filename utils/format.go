package utils

import (
	"encoding/json"
	"fmt"
)

var emptyStruct struct{}

const prettyPrintPrefix = ""
const prettyPrintIndent = "    "

func PrettyPrint(data interface{}) error {
	if data == nil {
		data = emptyStruct
	}
	jsonBytes, err := json.MarshalIndent(data, prettyPrintPrefix, prettyPrintIndent)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonBytes))
	return nil
}
