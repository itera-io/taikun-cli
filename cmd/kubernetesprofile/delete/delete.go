package delete

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/kubernetes_profiles"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <kubernetes-profile-id>...",
		Short: "Delete one or more kubernetes profiles",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ids, err := cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return cmdutils.DeleteMultiple(ids, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return cmd
}

func deleteRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := kubernetes_profiles.NewKubernetesProfilesDeleteParams().WithV(apiconfig.Version).WithID(id)
	_, _, err = apiClient.Client.KubernetesProfiles.KubernetesProfilesDelete(params, apiClient)
	if err == nil {
		format.PrintDeleteSuccess("Kubernetes profile", id)
	}

	return
}
