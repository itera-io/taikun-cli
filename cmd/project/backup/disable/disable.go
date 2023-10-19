package disable

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
	myApiClient := tk.NewClient()
	body := taikuncore.DisableBackupCommand{
		ProjectId: &opts.ProjectID,
	}
	_, err = myApiClient.Client.BackupPolicyAPI.BackupDisableBackup(context.TODO()).DisableBackupCommand(body).Execute()
	if err != nil {
		//return tk.CreateError(response, err)
		return cmderr.ErrProjectBackupAlreadyDisabled
	}
	out.PrintStandardSuccess()
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.DisableBackupCommand{
			ProjectID: opts.ProjectID,
		}

		params := backup.NewBackupDisableBackupParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.Backup.BackupDisableBackup(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}

/*
func getBackupCredentialID(projectID int32) (id int32, err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := servers.NewServersDetailsParams().WithV(taikungoclient.Version)
	params = params.WithProjectID(projectID)

	response, err := apiClient.Client.Servers.ServersDetails(params, apiClient)
	if err == nil {
		id = response.Payload.Project.S3CredentialID
		if id == 0 {
			err = cmderr.ErrProjectBackupAlreadyDisabled
		}
	}

	return
}
*/
