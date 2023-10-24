package list

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"METADATA-NAME", "metadataName",
		),
		field.NewVisible(
			"PROVIDER", "provider",
		),
		field.NewVisible(
			"NAMESPACE", "namespace",
		),
		field.NewVisible(
			"ACCESS-MODE", "accessMode",
		),
		field.NewVisible(
			"PHASE", "phase",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHiddenWithToStringFunc(
			"LAST-VALIDATED", "lastValidated", out.FormatDateTimeString,
		),
	},
)

type ListOptions struct {
	ProjectID int32
	Limit     int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <project-id>",
		Short: "List a project's backup sources",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddSortByAndReverseFlags(cmd, "backup-sources", listFields)
	cmdutils.AddColumnsFlag(cmd, listFields)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.BackupPolicyAPI.BackupListAllBackupStorages(context.TODO(), opts.ProjectID)
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var backupRestores = make([]taikuncore.BackupStorageLocationDto, 0)
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		backupRestores = append(backupRestores, data.GetData()...)

		count := int32(len(backupRestores))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}
		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(backupRestores)) > opts.Limit {
		backupRestores = backupRestores[:opts.Limit]
	}

	return out.PrintResults(backupRestores, listFields)

}
