package unbind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
					return cmderr.ErrIDArgumentNotANumber
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
	myApiClient := tk.NewClient()
	body := taikuncore.UnbindFlavorFromProjectCommand{
		Ids: bindings,
	}
	response, err := myApiClient.Client.FlavorsAPI.FlavorsUnbindFromProject(context.TODO()).UnbindFlavorFromProjectCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.UnbindFlavorFromProjectCommand{
			Ids: bindings,
		}
		params := flavors.NewFlavorsUnbindFromProjectParams().WithV(taikungoclient.Version).WithBody(&body)

		_, err = apiClient.Client.Flavors.FlavorsUnbindFromProject(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}
	*/

	return
}
