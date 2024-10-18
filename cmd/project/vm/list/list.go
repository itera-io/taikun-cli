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
		field.NewHidden(
			"HYPERVISOR", "hypervisor",
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

	cmdutils.AddSortByAndReverseFlags(&cmd, "standalone-project", listFields)
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

}
