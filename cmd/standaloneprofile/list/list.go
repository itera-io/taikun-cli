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
		Use:   "list",
		Short: "List standalone profiles",
		Args:  cobra.NoArgs, // FIXME maybe
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
