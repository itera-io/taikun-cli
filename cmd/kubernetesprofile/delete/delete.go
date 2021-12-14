package delete

import (
	"fmt"

	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/kubernetes_profiles"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <kubernetes-profile-id>",
		Short: "Delete a kubernetes profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cmdutils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given id must be a number")
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

	params := kubernetes_profiles.NewKubernetesProfilesDeleteParams().WithV(cmdutils.ApiVersion).WithID(id)
	_, _, err = apiClient.Client.KubernetesProfiles.KubernetesProfilesDelete(params, apiClient)
	if err == nil {
		fmt.Println("Kubernetes Profile deleted")
	}

	return
}
