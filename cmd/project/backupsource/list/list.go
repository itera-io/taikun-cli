package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/backup"
	"github.com/itera-io/taikungoclient/models"
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
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := backup.NewBackupListAllBackupStoragesParams().WithV(taikungoclient.Version)
	params.WithProjectID(opts.ProjectID)

	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var backupRestores = make([]*models.BackupStorageLocationDto, 0)

	for {
		response, err := apiClient.Client.Backup.BackupListAllBackupStorages(params, apiClient)
		if err != nil {
			return err
		}

		backupRestores = append(backupRestores, response.Payload.Data...)

		count := int32(len(backupRestores))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == response.Payload.TotalCount {
			break
		}

		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(backupRestores)) > opts.Limit {
		backupRestores = backupRestores[:opts.Limit]
	}

	return out.PrintResults(backupRestores, listFields)
}
