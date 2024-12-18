package list

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewHidden("ID", "packageId"),
		field.NewVisible("NAME", "normalizedName"),
		field.NewVisibleWithToStringFunc("REPOSITORY", "repository", out.FormatRepoName),
		field.NewVisible("VERSION", "appVersion"),
		field.NewVisible("DESCRIPTION", "description"),
		field.NewHidden("STARS", "stars"),
		field.NewHidden("DEPRECATED", "deprecated"),
		field.NewHidden("SIGNED", "signed"),
	},
)

type ListOptions struct {
	Repository string
	Limit      int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List available application",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().StringVarP(&opts.Repository, "repository", "r", "", "Name of the helm repository containing the applications")
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(&cmd, listFields)
	cmdutils.AddSortByAndReverseFlags(&cmd, "applications", listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()

	myRequest := myApiClient.Client.PackageAPI.PackageList(context.TODO())
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}
	if opts.Repository != "" {
		myRequest = myRequest.FilterBy(opts.Repository)
	}

	var applications = make([]taikuncore.AvailablePackagesDto, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		applications = append(applications, data.GetData()...)
		applicationsCount := int32(len(applications))

		if opts.Limit != 0 && applicationsCount >= opts.Limit {
			break
		}

		if applicationsCount == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(applicationsCount)
	}

	if opts.Limit != 0 && int32(len(applications)) > opts.Limit {
		applications = applications[:opts.Limit]
	}

	return out.PrintResults(applications, listFields)

}
