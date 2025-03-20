package disable

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type DisableOptions struct {
	RepoName       string
	RepoOrg        string
	OrganizationID int32
}

func NewCmdDisable() *cobra.Command {
	var opts DisableOptions

	cmd := cobra.Command{
		Use:   "disable <NAME> <ORG>",
		Short: "Disable a repository. Specify repository name and repository organization name.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.RepoName = args[0]
			opts.RepoOrg = args[1]
			return disableRun(opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", -1, "Organization ID")

	return &cmd
}

func disableRun(opts DisableOptions) (err error) {
	myApiClient := tk.NewClient()

	// Get ID
	var foundId string
	foundId = ""

	// Try recommended
	data, response, err := myApiClient.Client.AppRepositoriesAPI.RepositoryRecommendedList(context.TODO()).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	for _, repo := range data {
		if repo.Name == opts.RepoName && repo.OrganizationName == opts.RepoOrg {
			foundId = *repo.RepositoryId.Get()
			break
		}
	}
	// Try public
	if foundId == "" {
		data2, response, err := myApiClient.Client.AppRepositoriesAPI.RepositoryAvailableList(context.TODO()).IsPrivate(false).Search(opts.RepoName).Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		for _, repo := range data2.Data {
			if repo.Name == opts.RepoName && repo.OrganizationName == opts.RepoOrg {
				foundId = *repo.RepositoryId.Get()
				break
			}
		}
	}

	// Try private
	if foundId == "" {
		data3, response, err := myApiClient.Client.AppRepositoriesAPI.RepositoryAvailableList(context.TODO()).IsPrivate(true).Search(opts.RepoName).Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		for _, repo := range data3.Data {
			if repo.Name == opts.RepoName && repo.OrganizationName == opts.RepoOrg {
				foundId = *repo.RepositoryId.Get()
				break
			}
		}
	}

	command := taikuncore.UnbindAppRepositoryCommand{
		Ids: []string{foundId},
	}

	if opts.OrganizationID != -1 {
		command.OrganizationId.Set(&opts.OrganizationID)
	}

	_, response, err = myApiClient.Client.AppRepositoriesAPI.RepositoryUnbind(context.TODO()).UnbindAppRepositoryCommand(command).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return
}
