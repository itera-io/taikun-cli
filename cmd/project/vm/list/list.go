package list

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
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
		field.NewHiddenWithToStringFunc(
			"TAGS", "standAloneMetaDatas", out.FormatVMTags,
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
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddSortByAndReverseFlags(&cmd, "standalone-vms", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	vms, err := ListVMs(opts)
	if err == nil {
		return out.PrintResults(vms, listFields)
	}

	return
}

func ListVMs(opts *ListOptions) (vms []taikuncore.StandaloneVmsListForDetailsDto, err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	myRequest := myApiClient.Client.StandaloneAPI.StandaloneDetails(context.TODO(), opts.ProjectID)
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	// Execute a query into the API + graceful exit
	data, response, err := myRequest.Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return nil, err
	}
	// Manipulate the gathered data
	vms = data.GetData()
	return vms, nil
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := stand_alone.NewStandAloneDetailsParams().WithV(taikungoclient.Version)
		params = params.WithProjectID(opts.ProjectID)

		if config.SortBy != "" {
			params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
		}

		response, err := apiClient.Client.StandAlone.StandAloneDetails(params, apiClient)
		if err == nil {
			vms = response.Payload.Data
		}

		return
	*/
}
