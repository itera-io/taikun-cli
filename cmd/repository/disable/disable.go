package disable

import (
	"context"
	"fmt"

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
		Use:   "disable <REPOSITORY-NAME> <ORGANIZATION-NAME>",
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
	recommendCommand := myApiClient.Client.AppRepositoriesAPI.RepositoryRecommendedList(context.TODO())
	if opts.OrganizationID != -1 {
		recommendCommand = recommendCommand.OrganizationId(opts.OrganizationID)
	}
	data, response, err := recommendCommand.Execute()
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
		publicCommand := myApiClient.Client.AppRepositoriesAPI.RepositoryAvailableList(context.TODO()).IsPrivate(false).Search(opts.RepoName)
		if opts.OrganizationID != -1 {
			publicCommand = publicCommand.OrganizationId(opts.OrganizationID)
		}
		data2, response, err := publicCommand.Execute()
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
		privateCommand := myApiClient.Client.AppRepositoriesAPI.RepositoryAvailableList(context.TODO()).IsPrivate(true).Search(opts.RepoName)
		if opts.OrganizationID != -1 {
			privateCommand = privateCommand.OrganizationId(opts.OrganizationID)
		}
		data3, response, err := privateCommand.Execute()
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

	if foundId == "" {
		return fmt.Errorf("repo with name %s and org %s not enabled so you cannot disable", opts.RepoName, opts.RepoOrg)
	}

	command := taikuncore.UnbindAppRepositoryCommand{
		Ids: []string{foundId},
	}

	if opts.OrganizationID != -1 {
		command.OrganizationId.Set(&opts.OrganizationID)
	}

	response, err = myApiClient.Client.AppRepositoriesAPI.RepositoryUnbind(context.TODO()).UnbindAppRepositoryCommand(command).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return
}
