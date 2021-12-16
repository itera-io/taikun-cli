package list

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/config"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/flavors"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit                int32
	ProjectID            int32
	ReverseSortDirection bool
	SortBy               string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <project-id>",
		Short: "List a project's bound flavors",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.WrongIDArgumentFormatError
			}
			if opts.Limit < 0 {
				return cmderr.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return cmderr.OutputFormatInvalidError
			}
			opts.ProjectID = projectID
			return listRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results")
	cmd.Flags().StringVarP(&opts.SortBy, "sort-by", "s", "", "Sort results by attribute value")

	return cmd
}

func printResults(flavors []*models.BoundFlavorsForProjectsListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(flavors)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(flavors))
		for i, flavor := range flavors {
			data[i] = flavor
		}
		format.PrettyPrintTable(data,
			"id",
			"name",
			"cpu",
			"isAws",
			"isAzure",
			"isOpenstack",
			"projectName",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := flavors.NewFlavorsGetSelectedFlavorsForProjectParams().WithV(apiconfig.Version)
	params = params.WithProjectID(&opts.ProjectID)
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	flavors := []*models.BoundFlavorsForProjectsListDto{}
	for {
		response, err := apiClient.Client.Flavors.FlavorsGetSelectedFlavorsForProject(params, apiClient)
		if err != nil {
			return err
		}
		flavors = append(flavors, response.Payload.Data...)
		flavorsCount := int32(len(flavors))

		if opts.Limit != 0 && flavorsCount >= opts.Limit {
			break
		}
		if flavorsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&flavorsCount)
	}

	if opts.Limit != 0 && int32(len(flavors)) > opts.Limit {
		flavors = flavors[:opts.Limit]
	}

	printResults(flavors)
	return
}
