package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/kube_config"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "serviceAccountName",
		),
		field.NewVisible(
			"USERNAME", "userName",
		),
		field.NewVisible(
			"USER-ROLE", "userRole",
		),
		field.NewVisible(
			"PROJECT", "projectName",
		),
		field.NewVisible(
			"ACCESSIBLE-FOR-ALL", "isAccessibleForAll",
		),
		field.NewVisible(
			"KUBECONFIG-ROLE", "kubeConfigRoleName",
		),
		field.NewVisible(
			"CREATED-BY", "createdBy",
		),
	},
	// TODO FORMAT???
	// TODO check json
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

	cmdutils.AddSortByAndReverseFlags(&cmd, "kube-configs", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := kube_config.NewKubeConfigListParams().WithV(api.Version)
	params = params.WithProjectID(&opts.ProjectID)

	response, err := apiClient.Client.KubeConfig.KubeConfigList(params, apiClient)
	if err == nil {
		out.PrintResults(response.Payload.Data, listFields)
	}

	return
}
