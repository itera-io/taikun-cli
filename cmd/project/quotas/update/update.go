package update

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/project_quotas"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type UpdateOptions struct {
	QuotaID  int32
	DiskSize int
	RAM      int
	CPU      int64
}

func NewCmdUpdate() *cobra.Command {
	var opts UpdateOptions

	cmd := &cobra.Command{
		Use:   "update <project-id>",
		Short: "Update a project quota",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.WrongIDArgumentFormatError
			}
			opts.QuotaID = id
			return updateRun(&opts)
		},
	}

	cmd.Flags().Int64VarP(&opts.CPU, "cpu", "c", -1, "Maximum CPU usage (unlimited by default)")
	cmd.Flags().IntVarP(&opts.DiskSize, "disk-size", "d", -1, "Maximum Disk Size in GBs (unlimited by default)")
	cmd.Flags().IntVarP(&opts.RAM, "ram", "r", -1, "Maximum RAM in GBs (unlimited by default)")

	return cmd
}

func updateRun(opts *UpdateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.ProjectQuotaUpdateDto{
		IsCPUUnlimited:      true,
		IsDiskSizeUnlimited: true,
		IsRAMUnlimited:      true,
	}

	if opts.CPU > 0 {
		body.IsCPUUnlimited = false
		body.CPU = opts.CPU
	}
	if opts.DiskSize > 0 {
		body.IsDiskSizeUnlimited = false
		body.DiskSize = types.GiBToB(opts.DiskSize)
	}
	if opts.RAM > 0 {
		body.IsRAMUnlimited = false
		body.RAM = types.GiBToB(opts.RAM)
	}

	params := project_quotas.NewProjectQuotasEditParams().WithV(apiconfig.Version).WithBody(body).WithQuotaID(opts.QuotaID)
	_, err = apiClient.Client.ProjectQuotas.ProjectQuotasEdit(params, apiClient)
	if err == nil {
		format.PrintStandardSuccess()
	}

	return
}
