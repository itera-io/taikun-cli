package add

import (
	"context"
	"fmt"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
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
		field.NewVisibleWithToStringFunc(
			"USER-ID", "userId", out.FormatID,
		),
		field.NewVisible(
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

type AddOptions struct {
	AccessScope string
	ProjectID   int32
	Name        string
	Role        string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <project-id>",
		Short: "Add a kubeconfig to a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			if err := cmdutils.CheckFlagValue("role", opts.Role, types.KubeconfigRoles); err != nil {
				return err
			}
			if err := cmdutils.CheckFlagValue("access-scope", opts.AccessScope, types.KubeconfigAccessScopes); err != nil {
				return err
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AccessScope, "access-scope", "a", "",
		fmt.Sprintf(
			"Who can use the kubeconfig: `%s` (%s), `%s` (%s) or `%s` (%s) (required)",
			types.KubeconfigAccessPersonal, "only you",
			types.KubeconfigAccessManagers, "managers only",
			types.KubeconfigAccessAll, "all users with access to this project",
		),
	)
	cmdutils.MarkFlagRequired(&cmd, "access-scope")
	cmdutils.SetFlagCompletionValues(&cmd, "access-scope", types.KubeconfigAccessScopes.Keys()...)

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "name")

	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role (required)")
	cmdutils.MarkFlagRequired(&cmd, "role")
	cmdutils.SetFlagCompletionValues(&cmd, "role", types.KubeconfigRoles.Keys()...)

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.CreateKubeConfigCommand{}
	body.SetName(opts.Name)
	body.SetProjectId(opts.ProjectID)
	body.SetIsAccessibleForAll(opts.AccessScope == types.KubeconfigAccessAll)
	body.SetIsAccessibleForManager(opts.AccessScope == types.KubeconfigAccessManagers)
	body.SetKubeConfigRoleId(types.GetKubeconfigRole(opts.Role))

	data, response, err := myApiClient.Client.KubeConfigAPI.KubeconfigCreate(context.TODO()).CreateKubeConfigCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data, addFields)

}
