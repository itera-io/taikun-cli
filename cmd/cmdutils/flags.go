package cmdutils

import (
	"log"
	"reflect"
	"strings"

	"github.com/itera-io/taikun-cli/utils/list"
	"github.com/spf13/cobra"
)

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

func RegisterStaticFlagCompletion(cmd *cobra.Command, flagName string, values ...string) {
	RegisterFlagCompletionFunc(cmd, flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return values, cobra.ShellCompDirectiveDefault
	})
}

func getStructFieldJsonTag(structType reflect.Type, i int) string {
	return structType.Field(i).Tag.Get("json")
}

func extractNameFromJsonTag(tag string) (name string) {
	if strings.Count(tag, ",") == 0 {
		name = tag
	} else {
		tokens := strings.Split(tag, ",")
		name = tokens[0]
	}
	return
}

func getStructJsonTags(s interface{}) []string {
	structType := reflect.ValueOf(s).Type()
	structFieldCount := structType.NumField()
	structFieldJsonTags := make([]string, structFieldCount)
	for i := 0; i < structFieldCount; i++ {
		tag := getStructFieldJsonTag(structType, i)
		structFieldJsonTags[i] = extractNameFromJsonTag(tag)
	}
	return structFieldJsonTags
}

const (
	sortByFlag            = "sort-by"
	sortByFlagShorthand   = "s"
	sortByFlagDefault     = ""
	sortByFlagDescription = "Sort results by attribute value"
)

func AddSortByFlag(cmd *cobra.Command, optionStore *string, resultStruct interface{}) {
	cmd.Flags().StringVarP(
		optionStore,
		sortByFlag,
		sortByFlagShorthand,
		sortByFlagDefault,
		sortByFlagDescription,
	)
	resultStructJsonTags := getStructJsonTags(resultStruct)
	RegisterStaticFlagCompletion(cmd, sortByFlag, resultStructJsonTags...)
}

func AddIdOnlyFlag(cmd *cobra.Command, idOnlyFlagValue *bool) {
	cmd.Flags().BoolVarP(
		idOnlyFlagValue,
		"id-only",
		"I",
		false,
		"Output only the ID of the newly created resource (takes priority over the --format flag)",
	)
}

func AddLimitFlag(cmd *cobra.Command) {
	cmd.Flags().Int32VarP(&list.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
}
