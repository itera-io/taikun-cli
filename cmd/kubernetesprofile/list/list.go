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
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"TAIKUN-LB", "taikunLBEnabled",
		),
		field.NewVisible(
			"OCTAVIA", "octaviaEnabled",
		),
		field.NewVisible(
			"BASTION-PROXY", "exposeNodePortOnBastion",
		),
		field.NewVisible(
			"CNI", "cni",
		),
		field.NewVisible(
			"SCHEDULE-ON-MASTER", "allowSchedulingOnMaster",
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
		field.NewVisible(
			"NVIDIA-GPU", "nvidiaGpuOperatorEnabled",
		),
		field.NewVisible(
			"WASM", "wasmEnabled",
		),
		field.NewVisible(
			"PROXMOX-STORAGE", "proxmoxStorage",
		),
	},
)

type ListOptions struct {
	OrganizationID int32
	Limit          int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List kubernetes profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "kubernetes-profiles", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.KubernetesProfilesAPI.KubernetesprofilesList(context.TODO())
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var kubernetesProfiles = make([]taikuncore.KubernetesProfilesListDto, 0)
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		kubernetesProfiles = append(kubernetesProfiles, data.GetData()...)

		count := int32(len(kubernetesProfiles))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(kubernetesProfiles)) > opts.Limit {
		kubernetesProfiles = kubernetesProfiles[:opts.Limit]
	}

	return out.PrintResults(kubernetesProfiles, listFields)

}
