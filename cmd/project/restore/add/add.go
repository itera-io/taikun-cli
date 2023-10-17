package add

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "projectId",
		),
		field.NewVisible(
			"BACKUP-NAME", "backupName",
		),
		field.NewVisible(
			"RESTORE-NAME", "restoreName",
		),
	},
)

type AddOptions struct {
	IncludeNamespaces []string
	ExcludeNamespaces []string
	RestoreName       string
	BackupName        string
	ProjectID         int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <project-id>",
		Short: "Add a project's backup restore",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.BackupName, "backup-name", "b", "", "Backup Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "backup-name")
	cmd.Flags().StringVarP(&opts.RestoreName, "restore-name", "r", "", "Restore Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "restore-name")

	cmd.Flags().StringSliceVarP(&opts.IncludeNamespaces, "include-namespaces", "i", []string{}, "Included Namespaces")
	cmd.Flags().StringSliceVarP(&opts.ExcludeNamespaces, "exclude-namespaces", "e", []string{}, "Excluded Namespaces")

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.RestoreBackupCommand{
		ProjectId:         &opts.ProjectID,
		BackupName:        *taikuncore.NewNullableString(&opts.BackupName),
		RestoreName:       *taikuncore.NewNullableString(&opts.RestoreName),
		IncludeNamespaces: opts.IncludeNamespaces,
		ExcludeNamespaces: opts.ExcludeNamespaces,
	}
	response, err := myApiClient.Client.BackupPolicyAPI.BackupRestoreBackup(context.TODO()).RestoreBackupCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	//out.PrintResult(response, addFields) // Probably will not work. Come back and #FIXME
	out.PrintStandardSuccess()
	return

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.RestoreBackupCommand{
			ProjectID:         opts.ProjectID,
			BackupName:        opts.BackupName,
			RestoreName:       opts.RestoreName,
			IncludeNamespaces: opts.IncludeNamespaces,
			ExcludeNamespaces: opts.ExcludeNamespaces,
		}

		params := backup.NewBackupRestoreBackupParams().WithV(taikungoclient.Version).WithBody(&body)

		response, err := apiClient.Client.Backup.BackupRestoreBackup(params, apiClient)
		if err == nil {
			return out.PrintResult(response.Payload, addFields)
		}

		return
	*/
}
