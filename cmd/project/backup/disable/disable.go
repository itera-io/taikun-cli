package disable

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/backup"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type DisableOptions struct {
	ProjectID int32
}

func NewCmdDisable() *cobra.Command {
	var opts DisableOptions

	cmd := cobra.Command{
		Use:   "disable <project-id>",
		Short: "Disable backup for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return disableRun(&opts)
		},
	}

	return &cmd
}

func disableRun(opts *DisableOptions) (err error) {
	backupCredentialID, err := getBackupCredentialID(opts.ProjectID)
	if err != nil {
		return
	}

	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.DisableBackupCommand{
		ProjectID:      opts.ProjectID,
		S3CredentialID: backupCredentialID,
	}

	params := backup.NewBackupDisableBackupParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Backup.BackupDisableBackup(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}

func getBackupCredentialID(projectID int32) (id int32, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := servers.NewServersDetailsParams().WithV(api.Version)
	params = params.WithProjectID(projectID)

	response, err := apiClient.Client.Servers.ServersDetails(params, apiClient)
	if err == nil {
		id = response.Payload.Project.S3CredentialID
		if id == 0 {
			err = cmderr.ProjectBackupAlreadyDisabledError
		}
	}

	return
}
