package quota

import (
	"github.com/itera-io/taikun-cli/cmd/project/quota/edit"
	"github.com/itera-io/taikun-cli/cmd/project/quota/list"

	"github.com/spf13/cobra"
)

func NewCmdQuota() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quota <command>",
		Short: "Manage project quotas",
	}

	cmd.AddCommand(edit.NewCmdEdit())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
