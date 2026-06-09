package remove

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type DeleteOption struct {
	ProjectID int32
	Name      string
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOption
	cmd := cobra.Command{
		Use:   "delete <project-id>",
		Short: "Delete a project's backup",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			ctx, cancel := cmdutils.APIContext(cmd)
			defer cancel()
			return deleteRun(ctx, opts)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().StringVarP(&opts.Name, "backup-name", "", "", "Backup Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "backup-name")

	return &cmd
}

func deleteRun(ctx context.Context, opts DeleteOption) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.DeleteBackupCommand{
		ProjectId: &opts.ProjectID,
		Name:      *taikuncore.NewNullableString(&opts.Name),
	}
	response, err := myApiClient.Client.BackupPolicyAPI.BackupDeleteBackup(ctx).DeleteBackupCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintDeleteSuccess("Backup", opts.Name)
	return

}
