package importrepo

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type importOpts struct {
	Username     string
	Password     string
	Organization int32
	Url          string
}

func NewCmdEnable() *cobra.Command {
	var opts importOpts

	cmd := cobra.Command{
		Use:   "import <NAME>",
		Short: "Import a repository. Specify repository name and a URL and import it.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//repo, err :=args[0]
			return enableRun(args[0], opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.Organization, "organization", "o", 0, "Id of the organization to use for the list-public (partner)")
	_ = cmd.MarkFlagRequired("organization")

	cmd.Flags().StringVarP(&opts.Url, "url", "u", "", "URL to use for the import")
	_ = cmd.MarkFlagRequired("url")

	cmd.Flags().StringVarP(&opts.Username, "username", "n", "", "Username to use for the import")
	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "Password to use for the import")
	cmd.MarkFlagsRequiredTogether("username", "password")

	return &cmd
}

func enableRun(name string, opts importOpts) (err error) {
	myApiClient := tk.NewClient()

	//command := taikuncore.BindAppRepositoryCommand{
	//	FilteringElements: []taikuncore.FilteringElementDto{
	//		{
	//			OrganizationName: *taikuncore.NewNullableString(&org),
	//			Name:             *taikuncore.NewNullableString(&name),
	//		},
	//	},
	//}

	command := taikuncore.ImportRepoCommand{
		Name:           *taikuncore.NewNullableString(&name),
		Url:            *taikuncore.NewNullableString(&opts.Url),
		OrganizationId: *taikuncore.NewNullableInt32(&opts.Organization),
	}

	if opts.Username != "" {
		command.SetUsername(opts.Username)
	}
	if opts.Password != "" {
		command.SetPassword(opts.Password)
	}

	_, response, err := myApiClient.Client.AppRepositoriesAPI.RepositoryImport(context.TODO()).ImportRepoCommand(command).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return
}
