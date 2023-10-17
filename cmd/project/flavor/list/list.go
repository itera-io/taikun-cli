package list

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
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
	Limit     int32
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
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.ProjectID = projectID
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "bound-flavors", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.FlavorsAPI.FlavorsSelectedFlavorsForProject(context.TODO()).ProjectId(opts.ProjectID)
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	flavors := make([]taikuncore.BoundFlavorsForProjectsListDto, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		flavors = append(flavors, data.GetData()...)
		flavorsCount := int32(len(flavors))

		if opts.Limit != 0 && flavorsCount >= opts.Limit {
			break
		}

		if flavorsCount == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(flavorsCount)
	}

	if opts.Limit != 0 && int32(len(flavors)) > opts.Limit {
		flavors = flavors[:opts.Limit]
	}

	return out.PrintResults(flavors, listFields)

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := flavors.NewFlavorsGetSelectedFlavorsForProjectParams().WithV(taikungoclient.Version)
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

			if opts.Limit != 0 && flavorsCount >= opts.Limit {
				break
			}

			if flavorsCount == response.Payload.TotalCount {
				break
			}

			params = params.WithOffset(&flavorsCount)
		}

		if opts.Limit != 0 && int32(len(flavors)) > opts.Limit {
			flavors = flavors[:opts.Limit]
		}

		return out.PrintResults(flavors, listFields)
	*/
}
