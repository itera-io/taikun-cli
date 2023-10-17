package cmdutils

import (
	"context"
	"fmt"
	tk "github.com/Smidra/taikungoclient"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/gmap"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/spf13/cobra"
)

func MarkFlagRequired(cmd *cobra.Command, flag string) {
	if err := cmd.MarkFlagRequired(flag); err != nil {
		panic(err)
	}
}

func AddSortByAndReverseFlags(cmd *cobra.Command, sortType string, fields fields.Fields) {
	cmd.Flags().StringVarP(
		&config.SortBy,
		"sort-by",
		"S",
		"",
		"Sort results by attribute value",
	)
	SetFlagCompletionFunc(cmd, "sort-by", makeSortByCompletionFunc(sortType, fields))
	cmd.PreRunE = aggregateRunE(cmd.PreRunE, makeSortByPreRunE(fields))

	cmd.Flags().BoolVarP(
		&config.ReverseSortDirection,
		"reverse",
		"R",
		false,
		"Reverse order of results when passed with the --sort-by flag",
	)
}

func makeSortByPreRunE(fields fields.Fields) runE {
	return func(cmd *cobra.Command, args []string) error {
		if config.SortBy == "" {
			return nil
		}

		jsonPropertyName, found := fields.GetJsonPropertyNameFromName(config.SortBy)
		if !found {
			return fmt.Errorf("unknown sorting element '%s'", config.SortBy)
		}

		config.SortBy = jsonPropertyName

		return nil
	}
}

func makeSortByCompletionFunc(sortType string, fields fields.Fields) func(cmd *cobra.Command, args []string, toComplete string) []string {
	return func(cmd *cobra.Command, args []string, toComplete string) []string {
		sortingElements, err := getSortingElements(sortType)
		if err != nil {
			return []string{}
		}

		completions := make([]string, 0)

		for _, jsonPropertyName := range sortingElements {
			for _, field := range fields.AllFields() {
				if field.JsonPropertyName() == jsonPropertyName {
					completions = append(completions, field.Name())
					break
				}
			}
		}

		return lowerStringSlice(completions)
	}
}

func getSortingElements(sortType string) (sortingElements []string, err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.CommonAPI.CommonSortingElements(context.TODO(), sortType).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	sortingElements = data

	// Manipulate the gathered data
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := common.NewCommonGetSortingElementsParams().WithV(taikungoclient.Version)
		params = params.WithType(sortType)

		response, err := apiClient.Client.Common.CommonGetSortingElements(params, apiClient)
		if err != nil {
			return
		}

		sortingElements = response.Payload

		return
	*/
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
	SetFlagCompletionValues(cmd, "columns", lowerStringSlice(fields.AllNames())...)

	cmd.Flags().BoolVarP(
		&config.AllColumns,
		"all-columns",
		"A",
		false,
		"Display all columns in the output table (takes priority over the --columns flag)",
	)
}

func lowerStringSlice(stringSlice []string) []string {
	lower := make([]string, len(stringSlice))
	for i, str := range stringSlice {
		lower[i] = strings.ToLower(str)
	}

	return lower
}

func AddLimitFlag(cmd *cobra.Command, limit *int32) {
	cmd.Flags().Int32VarP(limit, "limit", "L", 0, "Limit number of results (limitless by default)")
	cmd.PreRunE = aggregateRunE(cmd.PreRunE,
		func(cmd *cobra.Command, args []string) error {
			if *limit < 0 {
				return cmderr.ErrNegativeLimit
			}
			return nil
		},
	)
}

func CheckFlagValue(flagName string, flagValue string, valid gmap.GenericMap) error {
	if !valid.Contains(flagValue) {
		return cmderr.UnknownFlagValueError(flagName, flagValue, valid.Keys())
	}

	return nil
}
