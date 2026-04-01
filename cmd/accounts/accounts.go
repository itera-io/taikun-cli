package accounts

import (
	addadmin "github.com/itera-io/taikun-cli/cmd/accounts/add-admin"
	"github.com/itera-io/taikun-cli/cmd/accounts/check"
	"github.com/itera-io/taikun-cli/cmd/accounts/create"
	"github.com/itera-io/taikun-cli/cmd/accounts/delete"
	"github.com/itera-io/taikun-cli/cmd/accounts/details"
	disable2fa "github.com/itera-io/taikun-cli/cmd/accounts/disable-2fa"
	enable2fa "github.com/itera-io/taikun-cli/cmd/accounts/enable-2fa"
	"github.com/itera-io/taikun-cli/cmd/accounts/groups"
	"github.com/itera-io/taikun-cli/cmd/accounts/list"
	"github.com/itera-io/taikun-cli/cmd/accounts/organizations"
	"github.com/itera-io/taikun-cli/cmd/accounts/projects"
	transferownership "github.com/itera-io/taikun-cli/cmd/accounts/transfer-ownership"
	"github.com/itera-io/taikun-cli/cmd/accounts/update"
	"github.com/itera-io/taikun-cli/cmd/accounts/users"
	"github.com/spf13/cobra"
)

func NewCmdAccounts() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "accounts  <command>",
		Short:   "Manage accounts in Taikun",
		Aliases: []string{"accounts", "acc", "A"},
	}

	cmd.AddCommand(create.NewCmdCreateAccount())
	cmd.AddCommand(check.NewCmdCheck())
	cmd.AddCommand(list.NewCmdListAccounts())
	cmd.AddCommand(delete.NewCmdDeleteAccount())
	cmd.AddCommand(details.NewCmdDetails())
	cmd.AddCommand(disable2fa.NewCmdDisable2fa())
	cmd.AddCommand(enable2fa.NewCmdEnable2fa())
	cmd.AddCommand(update.NewCmdUpdate())
	cmd.AddCommand(transferownership.NewCmdTransferOwnership())
	cmd.AddCommand(addadmin.NewCmdAddAdmin())
	cmd.AddCommand(organizations.NewCmdOrganizations())
	cmd.AddCommand(users.NewCmdUsers())
	cmd.AddCommand(projects.NewCmdProjects())
	cmd.AddCommand(groups.NewCmdGroups())
	//cmd.AddCommand(project.NewCmdCatalogProject())
	//cmd.AddCommand(makedefault.NewCmdMakedefault())
	//cmd.AddCommand(app.NewCmdApp())
	//cmd.AddCommand(lock.NewCmdLock())
	//cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
