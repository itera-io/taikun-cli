package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	ProjectID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <command>",
		Short: "List a project's servers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddSortByAndReverseFlags(&cmd, models.ServerListDto{})

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := servers.NewServersDetailsParams().WithV(apiconfig.Version)
	params = params.WithProjectID(opts.ProjectID)
	if config.SortBy != "" {
		if config.ReverseSortDirection {
			apiconfig.ReverseSortDirection()
		}
		params = params.WithSortBy(&config.SortBy)
		params = params.WithSortDirection(&apiconfig.SortDirection)
	}

	response, err := apiClient.Client.Servers.ServersDetails(params, apiClient)
	if err == nil {
		flavorNameField, err := getFlavorNameField(response.Payload.Data)
		if err != nil {
			return err
		}

		format.PrintResults(response.Payload.Data,
			"id",
			"name",
			"ipAdress",
			flavorNameField,
			"cpu",
			"ram",
			"diskSize",
			"role",
			"status",
			"createdAt",
		)
	}

	return
}

func getFlavorNameField(servers []*models.ServerListDto) (string, error) {
	if len(servers) == 0 {
		return "flavor", nil
	}
	if servers[0].AwsInstanceType != "" {
		return "awsInstanceType", nil
	}
	if servers[0].AzureVMSize != "" {
		return "azureVmSize", nil
	}
	if servers[0].OpenstackFlavor != "" {
		return "openstackFlavor", nil
	}
	return "", cmderr.ServerHasNoFlavorError
}
