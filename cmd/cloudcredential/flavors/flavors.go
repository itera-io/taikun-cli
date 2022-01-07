package flavors

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type FlavorsOptions struct {
	CloudCredentialID int32
	MaxCPU            int32
	MaxRAM            float64
	MinCPU            int32
	MinRAM            float64
}

func NewCmdFlavors() *cobra.Command {
	var opts FlavorsOptions

	cmd := &cobra.Command{
		Use:   "flavors <cloud-credential-id>",
		Short: "List a cloud credential's flavors",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cloudCredentialID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			opts.CloudCredentialID = cloudCredentialID
			return flavorRun(&opts)
		},
	}

	cmd.Flags().Int32Var(&opts.MaxCPU, "max-cpu", 36, "Maximal CPU count")
	cmd.Flags().Float64Var(&opts.MaxRAM, "max-ram", 500, "Maximal RAM size in GiB")
	cmd.Flags().Int32Var(&opts.MinCPU, "min-cpu", 2, "Minimal CPU count")
	cmd.Flags().Float64Var(&opts.MinRAM, "min-ram", 2, "Minimal RAM size in GiB")

	cmdutils.AddLimitFlag(cmd)
	cmdutils.AddSortByAndReverseFlags(cmd, models.FlavorsListDto{})

	return cmd
}

func flavorRun(opts *FlavorsOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := cloud_credentials.NewCloudCredentialsAllFlavorsParams().WithV(apiconfig.Version)
	params = params.WithCloudID(opts.CloudCredentialID)
	params = params.WithStartCPU(&opts.MinCPU).WithEndCPU(&opts.MaxCPU)
	minRAM := types.GiBToMiB(opts.MinRAM)
	maxRAM := types.GiBToMiB(opts.MaxRAM)
	params = params.WithStartRAM(&minRAM).WithEndRAM(&maxRAM)
	if config.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	flavors := []*models.FlavorsListDto{}
	for {
		response, err := apiClient.Client.CloudCredentials.CloudCredentialsAllFlavors(params, apiClient)
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

	format.PrintResults(flavors,
		"name",
		"cpu",
		"ram",
	)
	return
}
