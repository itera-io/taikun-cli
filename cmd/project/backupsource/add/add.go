package add

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "projectId",
		),
		field.NewVisible(
			"SOURCE-PROJECT-ID", "sourceProjectId",
		),
	},
)

type AddOptions struct {
	TargetProjectId int32
	SourceProjectId int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <project-id>",
		Short: "Add a project's backup source",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.TargetProjectId, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.SourceProjectId, "source-project-id", "s", 0, "Source Project Id (required)")
	cmdutils.MarkFlagRequired(&cmd, "source-project-id")

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	body := taikuncore.ImportBackupStorageLocationCommand{
		TargetProjectId: &opts.TargetProjectId,
		SourceProjectId: &opts.SourceProjectId,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.BackupPolicyAPI.BackupImportBackupStorage(context.TODO()).ImportBackupStorageLocationCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// out.PrintResult(response, addFields) // Probably will not work #FIXME
	//out.PrintStandardSuccess()
	return out.PrintResult(response, addFields)
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.ImportBackupStorageLocationCommand{
			TargetProjectID: opts.TargetProjectId,
			SourceProjectID: opts.SourceProjectId,
		}

		params := backup.NewBackupImportBackupStorageParams().WithV(taikungoclient.Version).WithBody(&body)

		response, err := apiClient.Client.Backup.BackupImportBackupStorage(params, apiClient)
		if err == nil {
			return out.PrintResult(response.Payload, addFields)
		}

		return
	*/
}
