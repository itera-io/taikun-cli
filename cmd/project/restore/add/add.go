package add

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

// New generated api client does not return data
//var addFields = fields.New(
//	[]*field.Field{
//		field.NewVisible(
//			"ID", "projectId",
//		),
//		field.NewVisible(
//			"BACKUP-NAME", "backupName",
//		),
//		field.NewVisible(
//			"RESTORE-NAME", "restoreName",
//		),
//	},
//)

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

	out.PrintStandardSuccess()
	return
}
