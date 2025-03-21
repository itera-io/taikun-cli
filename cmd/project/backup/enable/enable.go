package enable

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
	myApiClient := tk.NewClient()
	body := taikuncore.DeploymentEnableBackupCommand{
		ProjectId:      &opts.ProjectID,
		S3CredentialId: &opts.BackupCredentialID,
	}
	_, response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentEnableBackup(context.TODO()).DeploymentEnableBackupCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
