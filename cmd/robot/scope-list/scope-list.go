package scope_list

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible("KEY", "key"),
		field.NewVisible("TITLE", "title"),
		field.NewVisible("DESCRIPTION", "description"),
		field.NewVisible("TAG", "tag"),
	},
)

func NewCmdScopeList() *cobra.Command {
	cmd := cobra.Command{
		Use:   "scope-list",
		Short: "List available scopes for robot users",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRobots(cmd)
		},
	}

	return &cmd
}

func listRobots(cmd *cobra.Command) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	data, response, err := myApiClient.Client.RobotAPI.RobotScopeList(ctx).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResults(data, listFields)
}
