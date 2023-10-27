package showendpoints

import (
	"github.com/itera-io/taikun-cli/cmd/usertoken/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/spf13/cobra"
)

var infoFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"CONTROLLER", "controller",
		),
		field.NewVisible(
			"METHOD", "method",
		),
		field.NewVisible(
			"PATH", "path",
		),
		field.NewVisible(
			"DESCRIPTION", "description",
		),
	},
)

type ShowEndpointsOptions struct {
	usertokenName string
	showUnbound   bool
}

func NewCmdShowendpoints() *cobra.Command {
	var opts ShowEndpointsOptions

	cmd := cobra.Command{
		Use:   "show-endpoints <usertoken-name>",
		Short: "Show endpoints bound to a usertoken",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			infoFields.ShowAll()
			opts.usertokenName = args[0]
			return infoRun(&opts)
		},
	}

	complete.CompleteArgsWithUserTokenName(&cmd)
	cmd.Flags().BoolVarP(&opts.showUnbound, "unbound", "i", false, "Enable to show only endpoints not bound to the usertoken (inverse).")

	return &cmd
}

func infoRun(opts *ShowEndpointsOptions) (err error) {
	// Get tokenID of usertoken
	tokenID, err := complete.UserTokenIDFromUserTokenName(opts.usertokenName)
	if err != nil {
		return err
	}

	// Get endpoints for this token
	boundEndpoints, err := complete.GetAllBindingEndpoints(tokenID, opts.showUnbound)
	if err != nil {
		return err
	}

	return out.PrintResults(boundEndpoints, infoFields)
}
