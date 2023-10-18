package list

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var listFields = fields.NewNested(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"SIZE", "currentSize",
		),
		field.NewVisible(
			"TARGET-SIZE", "targetSize",
		),
		field.NewVisible(
			"VOLUME-TYPE", "volumeType",
		),
		field.NewVisible(
			"DEVICE-NAME", "deviceName",
		),
		field.NewVisible(
			"LUN-ID", "lunId",
		),
		field.NewVisible(
			"STATUS", "status",
		),
	},
	"disks",
)

type ListOptions struct {
	ProjectID      int32
	StandaloneVMID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <vm-id>",
		Short: "List a standalone VM's disks",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "project-id")

	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.StandaloneAPI.StandaloneDetails(context.TODO(), opts.ProjectID).Id(opts.StandaloneVMID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	return out.PrintResults(data.Data, listFields)
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := stand_alone.NewStandAloneDetailsParams().WithV(taikungoclient.Version)
		params = params.WithProjectID(opts.ProjectID).WithID(&opts.StandaloneVMID)

		response, err := apiClient.Client.StandAlone.StandAloneDetails(params, apiClient)
		if err == nil {
			return out.PrintResults(response.Payload.Data, listFields)
		}

		return
	*/
}
