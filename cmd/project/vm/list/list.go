package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/stand_alone"
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
			"FLAVOR", "currentFlavor",
		),
		field.NewVisible(
			"IP", "ipAddress",
		),
		field.NewVisible(
			"PUBLIC-IP", "publicIp",
		),
		field.NewVisible(
			"STATUS", "status",
		),
		field.NewVisible(
			"PROFILE", "profile/name",
		),
		field.NewHidden(
			"PROFILE-ID", "profile/id",
		),
		field.NewVisible(
			"IMAGE", "imageName",
		),
		field.NewHidden(
			"IMAGE-ID", "imageId",
		),
		field.NewHidden(
			"SSH-PUBLIC-KEY", "sshPublicKey",
		),
		field.NewHidden(
			"VOLUME-SIZE", "volumeSize",
		),
		field.NewHidden(
			"VOLUME-TYPE", "volumeType",
		),
		field.NewVisibleWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
		field.NewHiddenWithToStringFunc(
			"LAST-MODIFIED", "lastModified", out.FormatDateTimeString,
		),
		field.NewHidden(
			"LAST-MODIFIED-BY", "lastModifiedBy",
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
		Short: "List a project's standalone VMs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return listRun(&opts)
		},
	}

	cmdutils.AddSortByAndReverseFlags(&cmd, "standalone-vms", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	vms, err := ListVMs(opts)
	if err == nil {
		out.PrintResults(vms, listFields)
	}

	return
}

func ListVMs(opts *ListOptions) (vms []*models.StandaloneVmsListForDetailsDto, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := stand_alone.NewStandAloneDetailsParams().WithV(api.Version)
	params = params.WithProjectID(opts.ProjectID)
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	response, err := apiClient.Client.StandAlone.StandAloneDetails(params, apiClient)
	if err == nil {
		vms = response.Payload.Data
	}

	return
}
