package list_subnets

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "subnetId",
		),
		field.NewVisible(
			"TYPE", "subnetType",
		),
	},
)

func NewCmdListSubnets() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list-subnets",
		Short: "List subnets for a specific project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			projectID, err := types.Atoi32(args[0])
			if err != nil {
				return
			}
			return listSubnetsRun(projectID)
		},
		Aliases: cmdutils.ListAliases,
	}

	return &cmd
}

func listSubnetsRun(projectID int32) (err error) {
	myApiClient := tk.NewClient()
	req := myApiClient.Client.ServersAPI.ServersDetails(context.TODO(), projectID)
	details, response, err := req.Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResults(details.Project.CloudSubnets, listFields)

}
