package enable

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/backup"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type EnableOptions struct {
	ProjectID          int32
	BackupCredentialID int32
}

func NewCmdEnable() *cobra.Command {
	var opts EnableOptions

	cmd := cobra.Command{
		Use:   "enable <project-id>",
		Short: "Enable backup for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return enableRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(
		&opts.BackupCredentialID,
		"backup-credential-id", "b", 0,
		"Backup credential ID (required)",
	)
	cmdutils.MarkFlagRequired(&cmd, "backup-credential-id")

	return &cmd
}

func enableRun(opts *EnableOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.EnableBackupCommand{
		ProjectID:      opts.ProjectID,
		S3CredentialID: opts.BackupCredentialID,
	}

	params := backup.NewBackupEnableBackupParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Backup.BackupEnableBackup(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
