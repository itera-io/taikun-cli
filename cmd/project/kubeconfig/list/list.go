package list

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "displayName",
		),
		field.NewVisible(
			"PROJECT", "projectName",
		),
		field.NewVisible(
			"ROLE", "kubeConfigRoleName",
		),
		field.NewVisible(
			"ALL-HAVE-ACCESS", "isAccessibleForAll",
		),
		field.NewVisible(
			"MANAGERS-HAVE-ACCESS", "isAccessibleForManager",
		),
		field.NewVisible(
			"USERNAME", "userName",
		),
		field.NewHiddenWithToStringFunc(
			"USER-ID", "userId", out.FormatID,
		),
		field.NewHidden(
			"USER-ROLE", "userRole",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
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
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.KubeConfigAPI.KubeconfigList(context.TODO()).ProjectId(opts.ProjectID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return out.PrintResults(data.GetData(), listFields)

}
