package cmdutils

import (
	"log"
	"reflect"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/gmap"
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

func frequencyMapFromStringSlice(stringSlice []string) map[string]int {
	freqMap := map[string]int{}
	for _, str := range stringSlice {
		freqMap[str] += 1
	}
	return freqMap
}

func GetCommonJsonTagsInStructs(structs []interface{}) []string {
	jsonTags := make([]string, 0)

	for _, s := range structs {
		jsonTags = append(jsonTags, getStructJsonTags(s)...)
	}

	jsonTagsFreqMap := frequencyMapFromStringSlice(jsonTags)

	structsCount := len(structs)
	commonJsonTags := make([]string, 0)
	for jsonTag, jsonTagFreq := range jsonTagsFreqMap {
		if jsonTagFreq == structsCount {
			commonJsonTags = append(commonJsonTags, jsonTag)
		}
	}

	return commonJsonTags
}

func AddSortByAndReverseFlags(cmd *cobra.Command, resultStructs ...interface{}) {
	cmd.Flags().StringVarP(
		&config.SortBy,
		"sort-by",
		"s",
		"",
		"Sort results by attribute value",
	)

	cmd.Flags().BoolVarP(
		&config.ReverseSortDirection,
		"reverse",
		"r",
		false,
		"Reverse order of results",
	)

	commonTags := GetCommonJsonTagsInStructs(resultStructs)

	RegisterStaticFlagCompletion(cmd, "sort-by", commonTags...)
}

func AddOutputOnlyIDFlag(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(
		&config.OutputOnlyID,
		"id-only",
		"I",
		false,
		"Output only the ID of the newly created resource (takes priority over the --format flag)",
	)
}

func AddLimitFlag(cmd *cobra.Command) {
	cmd.Flags().Int32VarP(&config.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
}

func CheckFlagValue(flagName string, flagValue string, valid gmap.GenericMap) error {
	if !valid.Contains(flagValue) {
		return cmderr.UnknownFlagValueError(flagName, flagValue, valid.Keys())
	}
	return nil
}
