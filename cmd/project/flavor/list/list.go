package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/flavors"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"CPU", "cpu",
		),
		field.NewVisibleWithToStringFunc(
			"RAM", "ram", out.FormatBToGiB,
		),
		field.NewVisible(
			"AWS", "isAws",
		),
		field.NewVisible(
			"AZURE", "isAzure",
		),
		field.NewVisible(
			"OPENSTACK", "isOpenstack",
		),
		field.NewHidden(
			"PROJECT", "projectName",
		),
		field.NewHidden(
			"PROJECT-ID", "projectId",
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
		Short: "List a project's bound flavors",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			opts.ProjectID = projectID
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(&cmd)
	cmdutils.AddSortByAndReverseFlags(&cmd, "bound-flavors", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := flavors.NewFlavorsGetSelectedFlavorsForProjectParams().WithV(api.Version)
	params = params.WithProjectID(&opts.ProjectID)
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	flavors := []*models.BoundFlavorsForProjectsListDto{}
	for {
		response, err := apiClient.Client.Flavors.FlavorsGetSelectedFlavorsForProject(params, apiClient)
		if err != nil {
			return err
		}
		flavors = append(flavors, response.Payload.Data...)
		flavorsCount := int32(len(flavors))

		if config.Limit != 0 && flavorsCount >= config.Limit {
			break
		}
		if flavorsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&flavorsCount)
	}

	if config.Limit != 0 && int32(len(flavors)) > config.Limit {
		flavors = flavors[:config.Limit]
	}

	out.PrintResults(flavors, listFields)
	return
}
