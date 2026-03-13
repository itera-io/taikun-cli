package disable

import (
	"context"
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type DisableOptions struct {
	RepoName       string
	OrganizationID int32
}

func NewCmdDisable() *cobra.Command {
	var opts DisableOptions

	cmd := cobra.Command{
		Use:   "disable <REPOSITORY-NAME>",
		Short: "Disable a repository. Specify repository name.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.RepoName = args[0]
			return disableRun(opts)
		},
	}

	cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)

	return &cmd
}

func disableRun(opts DisableOptions) (err error) {
	orgID, err := cmdutils.ResolveOrgID(opts.OrganizationID, cmdutils.IsRobotAuth())
	if err != nil {
		return err
	}

	myApiClient := tk.NewClient()

	// Get ID
	var foundId string
	foundId = ""

	// Try recommended
	recommendCommand := myApiClient.Client.AppRepositoriesAPI.RepositoryRecommendedList(context.TODO())
	if orgID != 0 {
		recommendCommand = recommendCommand.OrganizationId(orgID)
	}
	data, response, err := recommendCommand.Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	for _, repo := range data {
		if repo.Name == opts.RepoName {
			foundId = *repo.RepositoryId.Get()
			break
		}
	}
	// Try public
	if foundId == "" {
		publicCommand := myApiClient.Client.AppRepositoriesAPI.RepositoryAvailableList(context.TODO()).IsPrivate(false).Search(opts.RepoName)
		if orgID != 0 {
			publicCommand = publicCommand.OrganizationId(orgID)
		}
		data2, response, err := publicCommand.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		for _, repo := range data2.Data {
			if repo.Name == opts.RepoName {
				foundId = *repo.RepositoryId.Get()
				break
			}
		}
	}

	// Try private
	if foundId == "" {
		privateCommand := myApiClient.Client.AppRepositoriesAPI.RepositoryAvailableList(context.TODO()).IsPrivate(true).Search(opts.RepoName)
		if orgID != 0 {
			privateCommand = privateCommand.OrganizationId(orgID)
		}
		data3, response, err := privateCommand.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		for _, repo := range data3.Data {
			if repo.Name == opts.RepoName {
				foundId = *repo.RepositoryId.Get()
				break
			}
		}
	}

	if foundId == "" {
		return fmt.Errorf("repo with name %s not enabled so you cannot disable", opts.RepoName)
	}

	command := taikuncore.UnbindAppRepositoryCommand{
		Ids: []string{foundId},
	}

	if orgID != 0 {
		command.OrganizationId = &orgID
	}

	response, err = myApiClient.Client.AppRepositoriesAPI.RepositoryUnbind(context.TODO()).UnbindAppRepositoryCommand(command).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return
}
