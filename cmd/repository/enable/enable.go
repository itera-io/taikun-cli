package enable

import (
	"context"
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type EnableOptions struct {
	RepoName       string
	OrganizationID int32
}

func NewCmdEnable() *cobra.Command {
	var opts EnableOptions

	cmd := cobra.Command{
		Use:   "enable <REPOSITORY-NAME>",
		Short: "Enable a repository. Specify repository name.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.RepoName = args[0]
			return enableRun(opts)
		},
	}

	cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)

	return &cmd
}

func enableRun(opts EnableOptions) (err error) {
	orgID, err := cmdutils.ResolveOrgID(opts.OrganizationID, cmdutils.IsRobotAuth())
	if err != nil {
		return err
	}

	orgIDStr := fmt.Sprintf("%s", orgID)

	myApiClient := tk.NewClient()

	command := taikuncore.BindAppRepositoryCommand{
		FilteringElements: []taikuncore.FilteringElementDto{
			{
				Name:             *taikuncore.NewNullableString(&opts.RepoName),
				OrganizationName: *taikuncore.NewNullableString(&orgIDStr),
			},
		},
	}
	if orgID != 0 {
		command.OrganizationId = &orgID
	}

	response, err := myApiClient.Client.AppRepositoriesAPI.RepositoryBind(context.TODO()).BindAppRepositoryCommand(command).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return
}
