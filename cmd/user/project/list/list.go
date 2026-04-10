package list

import (
	"context"
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/user/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "projectId",
		),
		field.NewVisible(
			"NAME", "projectName",
		),
		field.NewVisible(
			"IS-BOUND", "isBound",
		),
		field.NewVisible(
			"KUBERNETES-VERSION", "kubernetesVersion",
		),
		field.NewVisible(
			"IS-LOCKED", "isLocked",
		),
		field.NewVisible(
			"MAINTENANCE-MODE-ENABLED", "maintenanceModeEnabled",
		),
		field.NewVisible(
			"VIRTUAL-CLUSTER", "isVirtualCluster",
		),
		field.NewVisible(
			"STATUS", "status",
		),
		field.NewVisible(
			"HEALTH", "health",
		),
	},
)

type ListOptions struct {
	UserID string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <user-id>",
		Short: "List a user's assigned projects",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.UserID = args[0]
			return listRun(&opts)
		},
	}

	complete.CompleteArgsWithUserID(&cmd)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

// listRun calls the API, gets the Users and prints their bound projects.
func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.ProjectsAPI.ProjectsDropdown(context.TODO()).UserId(opts.UserID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	// No projects found for user.
	if len(data) == 0 {
		return fmt.Errorf("no projects found for user with ID %s", opts.UserID)
	}
	return out.PrintResults(data, listFields)
}
