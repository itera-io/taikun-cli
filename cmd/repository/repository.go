package repository

import (
	deleterepo "github.com/itera-io/taikun-cli/cmd/repository/delete"
	"github.com/itera-io/taikun-cli/cmd/repository/disable"
	"github.com/itera-io/taikun-cli/cmd/repository/enable"
	importrepo "github.com/itera-io/taikun-cli/cmd/repository/import"
	list_private "github.com/itera-io/taikun-cli/cmd/repository/list-private"
	list_public "github.com/itera-io/taikun-cli/cmd/repository/list-public"
	list_recommend "github.com/itera-io/taikun-cli/cmd/repository/list-recommend"
	"github.com/spf13/cobra"
)

func NewCmdRepository() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "repository <command>",
		Short:   "Explore available repositories in Taikun",
		Aliases: []string{"repositories", "repo", "repos"},
	}

	cmd.AddCommand(list_recommend.NewCmdList())
	cmd.AddCommand(list_public.NewCmdList())
	cmd.AddCommand(list_private.NewCmdList())
	cmd.AddCommand(enable.NewCmdEnable())
	cmd.AddCommand(disable.NewCmdDisable())
	cmd.AddCommand(importrepo.NewCmdEnable())
	cmd.AddCommand(deleterepo.NewCmdDelete())

	return cmd
}
