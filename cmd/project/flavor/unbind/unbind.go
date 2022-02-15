package unbind

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
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
				binding, err := types.Atoi32(arg)
				if err != nil {
					return cmderr.IDArgumentNotANumberError
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
	params := flavors.NewFlavorsUnbindFromProjectParams().WithV(api.Version).WithBody(&body)
	_, err = apiClient.Client.Flavors.FlavorsUnbindFromProject(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
