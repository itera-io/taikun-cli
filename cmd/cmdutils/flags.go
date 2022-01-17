package cmdutils

import (
	"log"
	"reflect"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/gmap"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/spf13/cobra"
)

func MarkFlagRequired(cmd *cobra.Command, flag string) {
	if err := cmd.MarkFlagRequired(flag); err != nil {
		log.Fatal(err)
	}
}

type FlagCompCoreFunc func(cmd *cobra.Command, args []string, toComplete string) []string

func RegisterFlagCompletionFunc(cmd *cobra.Command, flagName string, f FlagCompCoreFunc) {
	if err := cmd.RegisterFlagCompletionFunc(flagName, makeFlagCompFunc(f)); err != nil {
		log.Fatal(err)
	}
}

func makeFlagCompFunc(f FlagCompCoreFunc) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return f(cmd, args, toComplete), cobra.ShellCompDirectiveNoFileComp
	}
}

func RegisterFlagCompletion(cmd *cobra.Command, flagName string, values ...string) {
	RegisterFlagCompletionFunc(cmd, flagName, func(cmd *cobra.Command, args []string, toComplete string) []string {
		return values
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

func getCommonJsonTagsInStructs(structs []interface{}) []string {
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

func AddSortByAndReverseFlags(cmd *cobra.Command, fields fields.Fields) {
	cmd.Flags().StringVarP(
		&config.SortBy,
		"sort-by",
		"S",
		"",
		"Sort results by attribute value",
	)

	cmd.Flags().BoolVarP(
		&config.ReverseSortDirection,
		"reverse",
		"R",
		false,
		"Reverse order of results when passed with the --sort-by flag",
	)

	RegisterFlagCompletion(cmd, "sort-by", fields.AllNames()...)
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

func AddColumnsFlag(cmd *cobra.Command, fields fields.Fields) {
	cmd.Flags().StringSliceVarP(
		&config.Columns,
		"columns",
		"C",
		[]string{},
		"Specify which columns to display in the output table",
	)
	RegisterFlagCompletion(cmd, "columns", fields.AllNames()...)

	cmd.Flags().BoolVarP(
		&config.AllColumns,
		"all-columns",
		"A",
		false,
		"Display all columns in the output table (takes priority over the --columns flag)",
	)
}

func AddLimitFlag(cmd *cobra.Command) {
	cmd.Flags().Int32VarP(&config.Limit, "limit", "L", 0, "Limit number of results (limitless by default)")
}

func CheckFlagValue(flagName string, flagValue string, valid gmap.GenericMap) error {
	if !valid.Contains(flagValue) {
		return cmderr.UnknownFlagValueError(flagName, flagValue, valid.Keys())
	}
	return nil
}
