package add

import (
	"fmt"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/kube_config"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
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

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CreateKubeConfigCommand{
		IsAccessibleForAll:     opts.AccessScope == types.KubeconfigAccessAll,
		IsAccessibleForManager: opts.AccessScope == types.KubeconfigAccessManagers,
		KubeConfigRoleID:       types.GetKubeconfigRole(opts.Role),
		Name:                   opts.Name,
		ProjectID:              opts.ProjectID,
	}

	params := kube_config.NewKubeConfigCreateParams().WithV(api.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.KubeConfig.KubeConfigCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"serviceAccountName",
			"userName",
			"userRole",
			"projectName",
			"isAccessibleForAll",
			"kubeConfigRoleName",
		)
	}

	return
}
