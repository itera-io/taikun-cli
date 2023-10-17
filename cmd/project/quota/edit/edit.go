package edit

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type EditOptions struct {
	projectId      int32
	serverCpu      int64
	serverRam      int
	serverDiskSize int
	vmCpu          int64
	vmRam          int
	vmVolumeSize   int64
}

func NewCmdEdit() *cobra.Command {
	var opts EditOptions

	cmd := &cobra.Command{
		Use:   "edit <project-id>",
		Short: "Edit a project quota",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.projectId = id
			return editRun(&opts)
		},
	}

	cmd.Flags().Int64VarP(&opts.serverCpu, "server-cpu", "c", 1000000, "Maximum CPU for servers (unlimited by default)")
	cmd.Flags().IntVarP(&opts.serverDiskSize, "disk-size", "d", 102400, "Maximum Disk Size for servers in GBs (unlimited by default)")
	cmd.Flags().IntVarP(&opts.serverRam, "server-ram", "r", 102400, "Maximum RAM for servers in GBs (unlimited by default)")
	cmd.Flags().Int64VarP(&opts.vmCpu, "vm-cpu", "p", 1000000, "Maximum CPU for virtual machines (unlimited by default)")
	cmd.Flags().Int64VarP(&opts.vmVolumeSize, "vm-volume-size", "v", 102400, "Maximum Volume Size for virtual machines in GBs (unlimited by default)")
	cmd.Flags().IntVarP(&opts.vmRam, "vm-ram", "a", 102400, "Maximum RAM for virtual machines in GBs (unlimited by default)")

	return cmd
}

func editRun(opts *EditOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.UpdateQuotaCommand{
		QuotaId: &opts.projectId,
	}
	if opts.serverCpu > 0 {
		body.SetServerCpu(opts.serverCpu)
	}
	if opts.serverDiskSize > 0 {
		body.SetServerDiskSize(types.GiBToB(opts.serverDiskSize))
	}
	if opts.serverRam > 0 {
		body.SetServerRam(types.GiBToB(opts.serverRam))
	}
	if opts.vmCpu > 0 {
		body.SetVmCpu(opts.vmCpu)
	}
	if opts.vmVolumeSize > 0 {
		body.SetVmVolumeSize(opts.vmVolumeSize)
	}
	if opts.vmRam > 0 {
		body.SetVmRam(types.GiBToB(opts.vmRam))
	}
	response, err := myApiClient.Client.ProjectQuotasAPI.ProjectquotasUpdate(context.TODO()).UpdateQuotaCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

	/*
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
	*/
}
