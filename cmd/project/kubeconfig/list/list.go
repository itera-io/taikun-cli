package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/kube_config"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	ProjectID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <project-id>",
		Short: "List a project's kubeconfigs",
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

	cmdutils.AddSortByAndReverseFlags(&cmd, models.KubeConfigForUserDto{})

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := kube_config.NewKubeConfigListParams().WithV(apiconfig.Version)
	params = params.WithProjectID(&opts.ProjectID)

	response, err := apiClient.Client.KubeConfig.KubeConfigList(params, apiClient)
	if err == nil {
		format.PrintResults(response.Payload.Data,
			"id",
			"serviceAccountName",
			"userName",
			"userRole",
			"projectName",
			"isAccessibleForAll",
			"kubeConfigRoleName",
			"createdBy",
		)
	}

	return
}
