package deleterepo

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	RepoName       string
	RepoOrg        string
	OrganizationID int32
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions
	cmd := cobra.Command{
		Use:   "delete <NAME> <ORG>",
		Short: "Delete a repository. Specify repository name and repository organization name.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			//repo, err :=args[0]
			opts.RepoName = args[0]
			opts.RepoOrg = args[1]
			return deleteRun(opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", -1, "Organization ID")

	return &cmd
}

func deleteRun(opts DeleteOptions) (err error) {
	myApiClient := tk.NewClient()

	// Get ID
	var foundId int32
	foundId = -1

	// Try private
	data3, response, err := myApiClient.Client.AppRepositoriesAPI.RepositoryAvailableList(context.TODO()).IsPrivate(true).Search(opts.RepoName).OrganizationId(opts.OrganizationID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	for _, repo := range data3.Data {
		if repo.Name == opts.RepoName && repo.OrganizationName == opts.RepoOrg {
			foundId = repo.AppRepoId
			break
		}
	}
	if foundId == -1 {
		return fmt.Errorf("repo with name %s and org %s not found", opts.RepoName, opts.RepoOrg)
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
