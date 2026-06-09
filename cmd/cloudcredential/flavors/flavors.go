package flavors

import (
	"fmt"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/utils"
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

var flavorsFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"CPU", "cpu",
		),
		field.NewVisibleWithToStringFunc(
			"RAM", "ram", out.FormatRAM,
		),
		field.NewHidden(
			"DESCRIPTION", "description",
		),
	},
)

type FlavorsOptions struct {
	CloudCredentialID int32
	CloudType         taikuncore.CloudType
	MaxCPU            int32
	MaxRAM            int32
	MinCPU            int32
	MinRAM            int32
	Limit             int32
}

func NewCmdFlavors() *cobra.Command {
	var opts FlavorsOptions

	cmd := cobra.Command{
		Use:   "flavors <cloud-credential-id>",
		Short: "List a cloud credential's flavors",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cloudCredentialID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.CloudCredentialID = cloudCredentialID

			cloudType, err := utils.GetCloudType(cmd, opts.CloudCredentialID)
			if err != nil {
				return err
			}
			opts.CloudType = cloudType

			return flavorRun(cmd, &opts)
		},
	}

	cmd.Flags().Int32Var(&opts.MinCPU, "min-cpu", 2, "Minimal CPU count")
	cmd.Flags().Int32Var(&opts.MaxCPU, "max-cpu", 36, "Maximal CPU count")
	cmd.Flags().Int32Var(&opts.MinRAM, "min-ram", 2, "Minimal RAM size in GiB")
	cmd.Flags().Int32Var(&opts.MaxRAM, "max-ram", 500, "Maximal RAM size in GiB")

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "flavors", flavorsFields)
	cmdutils.AddColumnsFlag(&cmd, flavorsFields)

	return &cmd
}

func flavorRun(cmd *cobra.Command, opts *FlavorsOptions) (err error) {
	switch opts.CloudType {
	case taikuncore.CLOUDTYPE_AWS:
		return getAwsFlavors(cmd, opts)
	case taikuncore.CLOUDTYPE_AZURE:
		return getAzureFlavors(cmd, opts)
	case taikuncore.CLOUDTYPE_PROXMOX:
		return getProxmoxFlavors(cmd, opts)
	case taikuncore.CLOUDTYPE_GOOGLE:
		return getGoogleFlavors(cmd, opts)
	case taikuncore.CLOUDTYPE_OPENSTACK:
		return getOpenstackFlavors(cmd, opts)
	case taikuncore.CLOUDTYPE_VSPHERE:
		return getVsphereFlavors(cmd, opts)
	default:
		return fmt.Errorf("could not determine cloud type")
	}
}

func getAwsFlavors(cmd *cobra.Command, opts *FlavorsOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.FlavorsAPI.FlavorsAwsInstanceTypes(ctx, opts.CloudCredentialID).StartCpu(opts.MinCPU).EndCpu(opts.MaxCPU)
	myRequest = myRequest.StartRam(types.GiBToB(opts.MinRAM) - 100000).EndRam(types.GiBToB(opts.MaxRAM) + 100000).Limit(1000)

	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var flavors = make([]taikuncore.AwsFlavorListDto, 0)
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

	return out.PrintResults(flavors, flavorsFields)
}

func getProxmoxFlavors(cmd *cobra.Command, opts *FlavorsOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.FlavorsAPI.FlavorsProxmoxFlavors(ctx, opts.CloudCredentialID).StartCpu(opts.MinCPU).EndCpu(opts.MaxCPU)
	myRequest = myRequest.StartRam(types.GiBToB(opts.MinRAM) - 100000).EndRam(types.GiBToB(opts.MaxRAM) + 100000)

	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var flavors = make([]taikuncore.ProxmoxFlavorData, 0)
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

	return out.PrintResults(flavors, flavorsFields)
}

func getOpenstackFlavors(cmd *cobra.Command, opts *FlavorsOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.FlavorsAPI.FlavorsOpenstackFlavors(ctx, opts.CloudCredentialID).StartCpu(opts.MinCPU).EndCpu(opts.MaxCPU)
	myRequest = myRequest.StartRam(types.GiBToB(opts.MinRAM) - 100000).EndRam(types.GiBToB(opts.MaxRAM) + 100000)

	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var flavors = make([]taikuncore.OpenstackFlavorListDto, 0)
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

	return out.PrintResults(flavors, flavorsFields)
}

func getAzureFlavors(cmd *cobra.Command, opts *FlavorsOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.FlavorsAPI.FlavorsAzureVmSizes(ctx, opts.CloudCredentialID).StartCpu(opts.MinCPU).EndCpu(opts.MaxCPU)
	myRequest = myRequest.StartRam(types.GiBToB(opts.MinRAM) - 100000).EndRam(types.GiBToB(opts.MaxRAM) + 100000)
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var flavors = make([]taikuncore.AzureFlavorsWithPriceDto, 0)
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

	return out.PrintResults(flavors, flavorsFields)
}

func getGoogleFlavors(cmd *cobra.Command, opts *FlavorsOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.FlavorsAPI.FlavorsGoogleMachineTypes(ctx, opts.CloudCredentialID).StartCpu(opts.MinCPU).EndCpu(opts.MaxCPU)
	myRequest = myRequest.StartRam(types.GiBToB(opts.MinRAM) - 100000).EndRam(types.GiBToB(opts.MaxRAM) + 100000)
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var flavors = make([]taikuncore.GoogleFlavorDto, 0)
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

	return out.PrintResults(flavors, flavorsFields)
}

func getVsphereFlavors(cmd *cobra.Command, opts *FlavorsOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.FlavorsAPI.FlavorsVsphereFlavors(ctx, opts.CloudCredentialID).StartCpu(opts.MinCPU).EndCpu(opts.MaxCPU)
	myRequest = myRequest.StartRam(types.GiBToB(opts.MinRAM) - 100000).EndRam(types.GiBToB(opts.MaxRAM) + 100000)
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var flavors = make([]taikuncore.VsphereFlavorData, 0)
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

	return out.PrintResults(flavors, flavorsFields)
}
