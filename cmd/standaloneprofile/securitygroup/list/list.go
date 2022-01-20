package list

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		// FIXME
	},
)

type ListOptions struct {
	// FIXME add options
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <standalone-profile-id>",
		Short: "List a standalone profile's security groups",
		Args:  cobra.ExactArgs(1), // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	// FIXME

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// FIXME
	return
}
