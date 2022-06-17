package add

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/backup"
	"github.com/itera-io/taikungoclient/models"
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
}
