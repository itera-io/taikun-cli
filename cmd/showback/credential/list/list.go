package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
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
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd)
	cmdutils.AddSortByFlag(&cmd, &opts.SortBy, models.ShowbackCredentialsListDto{})

	return &cmd
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
		if config.Limit != 0 && showbackCredentialsCount >= config.Limit {
			break
		}
		if showbackCredentialsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&showbackCredentialsCount)
	}

	if config.Limit != 0 && int32(len(showbackCredentials)) > config.Limit {
		showbackCredentials = showbackCredentials[:config.Limit]
	}

	format.PrintResults(showbackCredentials,
		"id",
		"name",
		"organizationName",
		"url",
		"createdAt",
		"isLocked",
	)
	return
}
