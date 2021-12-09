package cmdutils

import (
	"log"

	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

const ApiVersion = "1"

const prettyPrintPrefix = ""
const prettyPrintIndent = "    "

func PrettyPrint(data interface{}) error {
	jsonBytes, err := json.MarshalIndent(data, prettyPrintPrefix, prettyPrintIndent)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonBytes))
	return nil
}

func MarkFlagRequired(cmd *cobra.Command, flag string) {
	if err := cmd.MarkFlagRequired(flag); err != nil {
		log.Fatal(err)
	}
}
