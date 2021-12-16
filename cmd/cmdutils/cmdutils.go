package cmdutils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const ApiVersion = "1"

var SortDirection = "asc"

func ReverseSortDirection() {
	SortDirection = "desc"
}

const prettyPrintPrefix = ""
const prettyPrintIndent = "    "

var emptyStruct struct{}

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

func MarkFlagRequired(cmd *cobra.Command, flag string) {
	if err := cmd.MarkFlagRequired(flag); err != nil {
		log.Fatal(err)
	}
}

func RegisterFlagCompletionFunc(cmd *cobra.Command, flagName string, f func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)) {
	if err := cmd.RegisterFlagCompletionFunc(flagName, f); err != nil {
		log.Fatal(err)
	}
}
