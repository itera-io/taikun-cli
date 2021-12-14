package list

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/flavors"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	ProjectID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <project-id>",
		Short: "List a project's bound flavors",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, err := cmdutils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given ID must be a number")
			}
			opts.ProjectID = projectID
			return listRun(&opts)
		},
	}

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := flavors.NewFlavorsGetSelectedFlavorsForProjectParams().WithV(cmdutils.ApiVersion)
	params = params.WithProjectID(&opts.ProjectID)

	flavors := []*models.BoundFlavorsForProjectsListDto{}
	for {
		response, err := apiClient.Client.Flavors.FlavorsGetSelectedFlavorsForProject(params, apiClient)
		if err != nil {
			return err
		}

		flavors = append(flavors, response.Payload.Data...)
		flavorsCount := int32(len(flavors))
		if flavorsCount == response.Payload.TotalCount {
			break
		}
	}

	cmdutils.PrettyPrint(flavors)
	return
}
