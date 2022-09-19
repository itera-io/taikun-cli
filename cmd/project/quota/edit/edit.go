package edit

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/project_quotas"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type EditOptions struct {
	QuotaID  int32
	DiskSize int
	RAM      int
	CPU      int64
}

func NewCmdEdit() *cobra.Command {
	var opts EditOptions

	cmd := &cobra.Command{
		Use:   "edit <quota-id>",
		Short: "Edit a project quota",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.QuotaID = id
			return editRun(&opts)
		},
	}

	cmd.Flags().Int64VarP(&opts.CPU, "cpu", "c", -1, "Maximum CPU (unlimited by default)")
	cmd.Flags().IntVarP(&opts.DiskSize, "disk-size", "d", -1, "Maximum Disk Size in GBs (unlimited by default)")
	cmd.Flags().IntVarP(&opts.RAM, "ram", "r", -1, "Maximum RAM in GBs (unlimited by default)")

	return cmd
}

func editRun(opts *EditOptions) error {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return err
	}

	body := &models.UpdateQuotaCommand{
		QuotaID: opts.QuotaID,
	}

	if opts.CPU > 0 {
		body.ServerCPU = opts.CPU
	}

	if opts.DiskSize > 0 {
		body.ServerDiskSize = types.GiBToB(opts.DiskSize)
	}

	if opts.RAM > 0 {
		body.ServerRAM = types.GiBToB(opts.RAM)
	}

	params := project_quotas.NewProjectQuotasEditParams().WithV(taikungoclient.Version).WithBody(body)

	if _, err := apiClient.Client.ProjectQuotas.ProjectQuotasEdit(params, apiClient); err != nil {
		return err
	}

	out.PrintStandardSuccess()

	return nil
}
