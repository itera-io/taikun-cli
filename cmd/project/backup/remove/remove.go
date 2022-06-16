package remove

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/backup"
	"github.com/itera-io/taikungoclient/models"
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
			return deleteRun(opts)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().StringVarP(&opts.Name, "backup-name", "", "", "Backup Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "backup-name")

	return &cmd
}

func deleteRun(opts DeleteOption) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.DeleteBackupCommand{ProjectID: opts.ProjectID, Name: opts.Name}
	params := backup.NewBackupDeleteBackupParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Backup.BackupDeleteBackup(params, apiClient)
	if err == nil {
		out.PrintDeleteSuccess("Backup", opts.Name)
	}

	return
}
