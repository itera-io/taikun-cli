package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/list"

	"github.com/itera-io/taikungoclient/client/cloud_credentials"
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

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List OpenStack cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ListRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(cmd)
	cmdutils.AddSortByFlag(cmd, &opts.SortBy, models.OpenstackCredentialsListDto{})

	return cmd
}

func ListRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(apiconfig.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	var openstackCloudCredentials = make([]*models.OpenstackCredentialsListDto, 0)
	for {
		response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return err
		}
		openstackCloudCredentials = append(openstackCloudCredentials, response.Payload.Openstack...)
		count := int32(len(openstackCloudCredentials))
		if list.Limit != 0 && count >= list.Limit {
			break
		}
		if count == response.Payload.TotalCountOpenstack {
			break
		}
		params = params.WithOffset(&count)
	}

	if list.Limit != 0 && int32(len(openstackCloudCredentials)) > list.Limit {
		openstackCloudCredentials = openstackCloudCredentials[:list.Limit]
	}

	format.PrintResults(openstackCloudCredentials,
		"id",
		"name",
		"organizationName",
		"project",
		"user",
		"isDefault",
		"isLocked",
	)
	return
}
