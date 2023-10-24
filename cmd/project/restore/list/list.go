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
			"BACKUP-NAME", "backupName",
		),
		field.NewVisible(
			"SCHEDULE-NAME", "scheduleName",
		),
		field.NewVisible(
			"NAMESPACE", "namespace",
		),
		field.NewVisible(
			"PHASE", "phase",
		),
		field.NewHidden(
			"INCLUDE-NAMESPACES", "includeNamespaces",
		),
		field.NewHidden(
			"EXCLUDE-NAMESPACES", "excludeNamespaces",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHiddenWithToStringFunc(
			"START-TIME-STAMP", "startTimeStamp", out.FormatDateTimeString,
		),
		field.NewHiddenWithToStringFunc(
			"COMPLETION-TIME-STAMP", "completionDateTime", out.FormatDateTimeString,
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
		Short: "List a project's backup restores",
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

	cmdutils.AddSortByAndReverseFlags(cmd, "restores", listFields)
	cmdutils.AddColumnsFlag(cmd, listFields)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.BackupPolicyAPI.BackupListAllRestores(context.TODO(), opts.ProjectID)
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}
	var backupRestores = make([]taikuncore.CRestoreDto, 0)
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
