package unlock

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/kubernetes_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdUnlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <kubernetes-profile-id>",
		Short: "Unlock a kubernetes profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.WrongIDArgumentFormatError
			}
			return unlockRun(id)
		},
	}

	return cmd
}

func unlockRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.KubernetesProfilesLockManagerCommand{
		ID:   id,
		Mode: types.LockedMode,
	}
	params := kubernetes_profiles.NewKubernetesProfilesLockManagerParams().WithV(apiconfig.Version).WithBody(&body)
	_, err = apiClient.Client.KubernetesProfiles.KubernetesProfilesLockManager(params, apiClient)
	if err == nil {
		format.PrintStandardSuccess()
	}

	return
}
