package enable

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type EnableOptions struct {
	RepoName         string
	OrganizationID   int32
	OrganizationName string
}

func NewCmdEnable() *cobra.Command {
	var opts EnableOptions

	cmd := cobra.Command{
		Use:   "enable <REPOSITORY-NAME>",
		Short: "Enable a repository. Specify repository name.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.RepoName = args[0]
			return enableRun(cmd, opts)
		},
	}

	cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)
	cmd.Flags().StringVarP(&opts.OrganizationName, "orgname", "n", "", "Organization's Name")

	return &cmd
}

func enableRun(cmd *cobra.Command, opts EnableOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	orgID, err := cmdutils.ResolveOrgID(opts.OrganizationID, cmdutils.IsRobotAuth())
	if err != nil {
		return err
	}

	myApiClient := tk.NewClient()

	command := taikuncore.BindAppRepositoryCommand{
		FilteringElements: []taikuncore.FilteringElementDto{
			{
				Name:             *taikuncore.NewNullableString(&opts.RepoName),
				OrganizationName: *taikuncore.NewNullableString(&opts.OrganizationName),
			},
		},
	}
	if orgID != 0 {
		command.OrganizationId = *taikuncore.NewNullableInt32(&orgID)
	}

	response, err := myApiClient.Client.AppRepositoriesAPI.RepositoryBind(ctx).BindAppRepositoryCommand(command).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return
}
