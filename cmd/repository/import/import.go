package importrepo

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type importOpts struct {
	Username       string
	Password       string
	OrganizationID int32
	Url            string
}

func NewCmdEnable() *cobra.Command {
	var opts importOpts

	cmd := cobra.Command{
		Use:   "import <NAME>",
		Short: "Import a repository. Specify repository name and a URL and import it.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return enableRun(args[0], opts)
		},
	}

	cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)

	cmd.Flags().StringVarP(&opts.Url, "url", "u", "", "URL to use for the import")
	_ = cmd.MarkFlagRequired("url")

	cmd.Flags().StringVarP(&opts.Username, "username", "n", "", "Username to use for the import")
	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "Password to use for the import")
	cmd.MarkFlagsRequiredTogether("username", "password")

	return &cmd
}

func enableRun(name string, opts importOpts) (err error) {
	orgID, err := cmdutils.ResolveOrgID(opts.OrganizationID, cmdutils.IsRobotAuth())
	if err != nil {
		return err
	}

	myApiClient := tk.NewClient()

	command := taikuncore.ImportRepoCommand{
		Name: *taikuncore.NewNullableString(&name),
		Url:  *taikuncore.NewNullableString(&opts.Url),
	}
	if orgID != 0 {
		command.OrganizationId = &orgID
	}

	if opts.Username != "" {
		command.SetUsername(opts.Username)
	}
	if opts.Password != "" {
		command.SetPassword(opts.Password)
	}

	response, err := myApiClient.Client.AppRepositoriesAPI.RepositoryImport(context.TODO()).ImportRepoCommand(command).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return
}
