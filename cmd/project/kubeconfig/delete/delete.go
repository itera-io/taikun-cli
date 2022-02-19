package delete

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient/client/kube_config"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <kubeconfig-id>...",
		Short: "Delete one or more kubeconfigs",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ids, err := cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return cmdutils.DeleteMultiple(ids, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return &cmd
}

func deleteRun(kubeconfigID int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.DeleteKubeConfigCommand{ID: kubeconfigID}
	params := kube_config.NewKubeConfigDeleteParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.KubeConfig.KubeConfigDelete(params, apiClient)
	if err == nil {
		out.PrintDeleteSuccess("Kubeconfig", kubeconfigID)
	}

	return
}
