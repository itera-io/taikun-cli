package deleterepo

import (
	"context"
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	RepoName       string
	OrganizationID int32
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions
	cmd := cobra.Command{
		Use:   "delete <REPOSITORY-NAME>",
		Short: "Delete a private repository. Specify repository name.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.RepoName = args[0]
			return deleteRun(opts)
		},
	}

	cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)

	return &cmd
}

func deleteRun(opts DeleteOptions) (err error) {
	orgID, err := cmdutils.ResolveOrgID(opts.OrganizationID, cmdutils.IsRobotAuth())
	if err != nil {
		return err
	}

	myApiClient := tk.NewClient()

	// Get ID
	var foundId int32
	foundId = -1

	// Try private
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
			foundId = repo.AppRepoId
			break
		}
	}
	if foundId == -1 {
		return fmt.Errorf("repo with name %s not found", opts.RepoName)
	}

	command := taikuncore.DeleteRepositoryCommand{
		AppRepoId: &foundId,
	}

	response, err = myApiClient.Client.AppRepositoriesAPI.RepositoryDelete(context.TODO()).DeleteRepositoryCommand(command).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return
}
