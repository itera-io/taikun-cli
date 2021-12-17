package cmdutils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"taikun-cli/utils/types"

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

func ArgsToNumericalIDs(args []string) ([]int32, error) {
	ids := make([]int32, len(args))
	for i, arg := range args {
		id, err := types.Atoi32(arg)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}
	return ids, nil
}

type DeleteFunc func(int32) error

func DeleteMultiple(ids []int32, deleteFunc DeleteFunc) error {
	errorOccured := false
	for _, id := range ids {
		if err := deleteFunc(id); err != nil {
			fmt.Fprintln(os.Stderr, err)
			errorOccured = true
		}
	}
	if errorOccured {
		fmt.Println()
		return errors.New("Failed to delete one or more resources")
	}
	return nil
}
