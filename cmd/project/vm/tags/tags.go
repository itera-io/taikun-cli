package tags

import (
	"fmt"

	"github.com/spf13/cobra"
)

type TagsOptions struct {
	// FIXME
}

func NewCmdTags() *cobra.Command {
	var opts TagsOptions

	cmd := cobra.Command{
		Use:   "tags <vm-id>",
		Short: "List a standalone VM's tags",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// FIXME
			return tagsRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func tagsRun(opts *TagsOptions) (err error) {
	// FIXME
	fmt.Println("TODO")
	return
}
