package project

import (
	"github.com/itera-io/taikun-cli/cmd/project/backup"
	"github.com/itera-io/taikun-cli/cmd/project/commit"
	"github.com/itera-io/taikun-cli/cmd/project/create"
	"github.com/itera-io/taikun-cli/cmd/project/delete"
	"github.com/itera-io/taikun-cli/cmd/project/etc"
	"github.com/itera-io/taikun-cli/cmd/project/flavor"
	"github.com/itera-io/taikun-cli/cmd/project/info"
	"github.com/itera-io/taikun-cli/cmd/project/list"
	"github.com/itera-io/taikun-cli/cmd/project/lock"
	"github.com/itera-io/taikun-cli/cmd/project/quotas"
	"github.com/itera-io/taikun-cli/cmd/project/server"
	"github.com/itera-io/taikun-cli/cmd/project/unlock"

	"github.com/spf13/cobra"
)

func NewCmdProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project <command>",
		Short: "Manage projects",
	}

	cmd.AddCommand(backup.NewCmdBackup())
	cmd.AddCommand(commit.NewCmdCommit())
	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(etc.NewCmdEtc())
	cmd.AddCommand(flavor.NewCmdFlavor())
	cmd.AddCommand(info.NewCmdInfo())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(quotas.NewCmdQuotas())
	cmd.AddCommand(server.NewCmdServer())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
