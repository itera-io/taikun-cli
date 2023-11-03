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
//			"SOURCE-PROJECT-ID", "sourceProjectId",
//		),
//	},
//)

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

	out.PrintStandardSuccess()
	return
}
