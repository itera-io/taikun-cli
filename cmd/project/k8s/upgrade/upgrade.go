package upgrade

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

type UpgradeOptions struct {
	ProjectID int32
}

func NewCmdUpgrade() *cobra.Command {
	var opts UpgradeOptions

	cmd := cobra.Command{
		Use:   "upgrade <project-id>",
		Short: "Upgrade a project's version of Kubespray to the latest one",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return upgradeRun(&opts)
		},
	}

	return &cmd
}

func upgradeRun(opts *UpgradeOptions) (err error) {
	myApiClient := tk.NewClient()
	response, err := myApiClient.Client.ProjectsAPI.ProjectsUpgrade(context.TODO(), opts.ProjectID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
