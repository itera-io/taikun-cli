package delete

import (
	"taikun-cli/api"
	"taikun-cli/utils"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/kubernetes_profiles"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <kubernetes-profile-id>",
		Short: "Delete a kubernetes profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return utils.WrongIDArgumentFormatError
			}
			return deleteRun(id)
		},
	}

	return cmd
}

func deleteRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := kubernetes_profiles.NewKubernetesProfilesDeleteParams().WithV(utils.ApiVersion).WithID(id)
	_, _, err = apiClient.Client.KubernetesProfiles.KubernetesProfilesDelete(params, apiClient)
	if err == nil {
		utils.PrintDeleteSuccess("Kubernetes profile", id)
	}

	return
}
