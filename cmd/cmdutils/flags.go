package cmdutils

import (
	"log"
	"reflect"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils/options"
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

// TODO move to options package
func AddSortByAndReverseFlags(cmd *cobra.Command, opts options.ListSorter, fields fields.Fields) {
	cmd.Flags().StringVarP(
		opts.GetSortByOption(),
		"sort-by",
		"S",
		"",
		"Sort results by attribute value",
	)

	fieldNames := fields.AllNames()
	lowerStringSlice(fieldNames)
	RegisterFlagCompletion(cmd, "sort-by", fieldNames...)

	cmd.Flags().BoolVarP(
		opts.GetReverseSortDirectionOption(),
		"reverse",
		"R",
		false,
		"Reverse order of results when passed with the --sort-by flag",
	)
}

// TODO move to options package
func AddOutputOnlyIDFlag(cmd *cobra.Command, opts options.Creator) {
	cmd.Flags().BoolVarP(
		opts.GetOutputOnlyIDOption(),
		"id-only",
		"I",
		false,
		"Output only the ID of the newly created resource (takes priority over the --format flag)",
	)
}

// TODO move to options package
func AddColumnsFlag(cmd *cobra.Command, opts options.TableWriter, fields fields.Fields) {
	cmd.Flags().StringSliceVarP(
		opts.GetColumnsOption(),
		"columns",
		"C",
		[]string{},
		"Specify which columns to display in the output table",
	)
	columns := fields.AllNames()
	lowerStringSlice(columns)
	RegisterFlagCompletion(cmd, "columns", columns...)

	cmd.Flags().BoolVarP(
		opts.GetAllColumnsOption(),
		"all-columns",
		"A",
		false,
		"Display all columns in the output table (takes priority over the --columns flag)",
	)
}

func lowerStringSlice(stringSlice []string) {
	size := len(stringSlice)
	for i := 0; i < size; i++ {
		stringSlice[i] = strings.ToLower(stringSlice[i])
	}
}

// TODO move to options package
func AddLimitFlag(cmd *cobra.Command, opts options.ListLimiter) {
	cmd.Flags().Int32VarP(opts.GetLimitOption(), "limit", "L", 0, "Limit number of results (limitless by default)")
}

func CheckFlagValue(flagName string, flagValue string, valid gmap.GenericMap) error {
	if !valid.Contains(flagValue) {
		return cmderr.UnknownFlagValueError(flagName, flagValue, valid.Keys())
	}
	return nil
}
