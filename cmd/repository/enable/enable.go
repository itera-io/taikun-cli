package enable

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type EnableOptions struct {
	RepoName       string
	RepoOrg        string
	OrganizationID int32
}

func NewCmdEnable() *cobra.Command {
	var opts EnableOptions

	cmd := cobra.Command{
		Use:   "enable <REPOSITORY-NAME> <ORGANIZATION-NAME>",
		Short: "Enable a repository. Specify repository name and repository organization name.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.RepoName = args[0]
			opts.RepoOrg = args[1]
			return enableRun(opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", -1, "Organization ID")

	return &cmd
}

func enableRun(opts EnableOptions) (err error) {
	myApiClient := tk.NewClient()

	//data, response, err := myApiClient.Client.OrganizationsAPI.OrganizationsList(context.TODO()).Search(org).Execute()
	//if err != nil {
	//	return tk.CreateError(response, err)
	//}
	//var orgid int32 = -1
	//for _, organization := range data.Data {
	//	if *organization.Name.Get() == org {
	//		orgid = *organization.Id
	//		break
	//	}
	//}
	command := taikuncore.BindAppRepositoryCommand{
		FilteringElements: []taikuncore.FilteringElementDto{
			{
				OrganizationName: *taikuncore.NewNullableString(&opts.RepoOrg),
				Name:             *taikuncore.NewNullableString(&opts.RepoName),
			},
		},
		//OrganizationId: *taikuncore.NewNullableInt32(&orgid),
	}
	if opts.OrganizationID != -1 {
		command.OrganizationId.Set(&opts.OrganizationID)
	}

	response, err := myApiClient.Client.AppRepositoriesAPI.RepositoryBind(context.TODO()).BindAppRepositoryCommand(command).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return
}
