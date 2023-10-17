package remove

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type DeleteOption struct {
	ProjectID       int32
	StorageLocation string
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOption
	cmd := cobra.Command{
		Use:   "delete <project-id>",
		Short: "Delete a project's backup source",
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

	cmd.Flags().StringVarP(&opts.StorageLocation, "storage-location", "", "", "Storage Location (required)")
	cmdutils.MarkFlagRequired(&cmd, "storage-location")

	return &cmd
}

func deleteRun(opts DeleteOption) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	body := taikuncore.DeleteBackupStorageLocationCommand{
		ProjectId:       &opts.ProjectID,
		StorageLocation: *taikuncore.NewNullableString(&opts.StorageLocation),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.BackupPolicyAPI.BackupDeleteBackupLocation(context.TODO()).DeleteBackupStorageLocationCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("Backup Source", opts.StorageLocation)
	return

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.DeleteBackupStorageLocationCommand{ProjectID: opts.ProjectID, StorageLocation: opts.StorageLocation}
		params := backup.NewBackupDeleteBackupLocationParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.Backup.BackupDeleteBackupLocation(params, apiClient)
		if err == nil {
			out.PrintDeleteSuccess("Backup Source", opts.StorageLocation)
		}

		return
	*/
}
