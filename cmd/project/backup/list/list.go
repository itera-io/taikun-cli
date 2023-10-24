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
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"NAME", "metadataName",
		),
		field.NewVisible(
			"SCHEDULE", "scheduleName",
		),
		field.NewVisible(
			"LOCATION", "location",
		),
		field.NewVisible(
			"PHASE", "phase",
		),
		field.NewVisible(
			"NAMESPACE", "namespace",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHiddenWithToStringFunc(
			"EXPIRED-AT", "expiration", out.FormatDateTimeString,
		),
	},
)

type ListOptions struct {
	ProjectID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <project-id>",
		Short: "List a project's backups",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddSortByAndReverseFlags(&cmd, "backups", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.BackupPolicyAPI.BackupListAllBackups(context.TODO(), opts.ProjectID)
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	data, response, err := myRequest.Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return out.PrintResults(data.GetData(), listFields) // Are you sure about missing pagination? #FIXME

}
