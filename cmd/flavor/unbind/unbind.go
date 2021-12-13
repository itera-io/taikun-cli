package unbind

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/flavors"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdUnbind() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unbind <flavor-bound-id>...",
		Short: "Unbind one or multiple flavors from a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			bindings := make([]int32, len(args))
			for i, arg := range args {
				binding, err := cmdutils.Atoi32(arg)
				if err != nil {
					return fmt.Errorf("the given IDs must be numbers")
				}
				bindings[i] = binding
			}
			return unbindRun(bindings)
		},
		Args: cobra.MinimumNArgs(1),
	}

	return cmd
}

func unbindRun(bindings []int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.UnbindFlavorFromProjectCommand{
		Ids: bindings,
	}
	params := flavors.NewFlavorsUnbindFromProjectParams().WithV(cmdutils.ApiVersion).WithBody(&body)
	response, err := apiClient.Client.Flavors.FlavorsUnbindFromProject(params, apiClient)
	if err == nil {
		cmdutils.PrettyPrint(response.Payload)
	}

	return
}
