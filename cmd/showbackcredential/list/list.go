package list

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/config"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit                int32
	OrganizationID       int32
	ReverseSortDirection bool
	SortBy               string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List showback credentials",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByFlag(&cmd, &opts.SortBy, models.ShowbackCredentialsListDto{})

	return &cmd
}

func printResults(showbackCredentials []*models.ShowbackCredentialsListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(showbackCredentials)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(showbackCredentials))
		for i, showbackCredential := range showbackCredentials {
			data[i] = showbackCredential
		}
		format.PrettyPrintTable(data,
			"id",
			"name",
			"organizationName",
			"url",
			"createdAt",
			"isLocked",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := showback.NewShowbackCredentialsListParams().WithV(apiconfig.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	var showbackCredentials = make([]*models.ShowbackCredentialsListDto, 0)
	for {
		response, err := apiClient.Client.Showback.ShowbackCredentialsList(params, apiClient)
		if err != nil {
			return err
		}
		showbackCredentials = append(showbackCredentials, response.Payload.Data...)
		showbackCredentialsCount := int32(len(showbackCredentials))
		if opts.Limit != 0 && showbackCredentialsCount >= opts.Limit {
			break
		}
		if showbackCredentialsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&showbackCredentialsCount)
	}

	if opts.Limit != 0 && int32(len(showbackCredentials)) > opts.Limit {
		showbackCredentials = showbackCredentials[:opts.Limit]
	}

	printResults(showbackCredentials)
	return
}
