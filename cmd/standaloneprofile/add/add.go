package add

import (
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		// FIXME
	},
)

type AddOptions struct {
	// FIXME add options
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a standalone profile",
		Args:  cobra.ExactArgs(1), // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return addRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// FIXME
	return
}
