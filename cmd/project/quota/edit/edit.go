package edit

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

const (
	minServerCPU = 2
	maxServerCPU = 1000000

	minServerDisk = 30
	maxServerDisk = 102400

	minServerRAM = 2
	maxServerRAM = 102400

	minVMCpu = 1
	maxVMCpu = 1000000

	minVMVolume = 1.0
	maxVMVolume = 102400.0

	minVMRam = 1
	maxVMRam = 102400
)

type EditOptions struct {
	projectId      int32
	serverCpu      int64
	serverRam      int32
	serverDiskSize int32
	vmCpu          int64
	vmRam          int32
	vmVolumeSize   float64
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
			return editRun(cmd, &opts)
		},
	}

	cmd.Flags().Int64VarP(&opts.serverCpu, "server-cpu", "c", 10, "Maximum CPU for servers (unlimited by default)")
	cmd.Flags().Int32VarP(&opts.serverDiskSize, "disk-size", "d", 40, "Maximum Disk Size for servers in GBs (unlimited by default)")
	cmd.Flags().Int32VarP(&opts.serverRam, "server-ram", "r", 10, "Maximum RAM for servers in GBs (unlimited by default)")
	cmd.Flags().Int64VarP(&opts.vmCpu, "vm-cpu", "p", 10, "Maximum CPU for virtual machines (unlimited by default)")
	cmd.Flags().Float64VarP(&opts.vmVolumeSize, "vm-volume-size", "v", 40, "Maximum Volume Size for virtual machines in GBs (unlimited by default)")
	cmd.Flags().Int32VarP(&opts.vmRam, "vm-ram", "a", 10, "Maximum RAM for virtual machines in GBs (unlimited by default)")

	return cmd
}

func editRun(cmd *cobra.Command, opts *EditOptions) error {
	client := tk.NewClient()

	// Load existing quotas
	data, response, err := client.Client.ProjectQuotasAPI.ProjectquotasList(context.TODO()).Id(opts.projectId).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	if data.GetTotalCount() != 1 {
		return fmt.Errorf("could not find quota for project with id %d", opts.projectId)
	}

	loaded := data.GetData()[0]

	// Start with loaded values
	serverCpu := loaded.ServerCpu
	serverDiskSize := int32(types.BToGiB(loaded.ServerDiskSize))
	serverRam := int32(types.BToGiB(loaded.ServerRam))
	vmCpu := loaded.VmCpu
	vmVolumeSize := loaded.VmVolumeSize
	vmRam := int32(types.BToGiB(loaded.VmRam))

	// Override with validated user input if flags are set
	if cmd.Flags().Changed("server-cpu") {
		if opts.serverCpu < minServerCPU || opts.serverCpu > maxServerCPU {
			return fmt.Errorf("server-cpu must be between %d and %d", minServerCPU, maxServerCPU)
		}
		serverCpu = opts.serverCpu
	}

	if cmd.Flags().Changed("disk-size") {
		if opts.serverDiskSize < minServerDisk || opts.serverDiskSize > maxServerDisk {
			return fmt.Errorf("disk-size must be between %d and %d", minServerDisk, maxServerDisk)
		}
		serverDiskSize = opts.serverDiskSize
	}

	if cmd.Flags().Changed("server-ram") {
		if opts.serverRam < minServerRAM || opts.serverRam > maxServerRAM {
			return fmt.Errorf("server-ram must be between %d and %d", minServerRAM, maxServerRAM)
		}
		serverRam = opts.serverRam
	}

	if cmd.Flags().Changed("vm-cpu") {
		if opts.vmCpu < minVMCpu || opts.vmCpu > maxVMCpu {
			return fmt.Errorf("vm-cpu must be between %d and %d", minVMCpu, maxVMCpu)
		}
		vmCpu = opts.vmCpu
	}

	if cmd.Flags().Changed("vm-volume-size") {
		if opts.vmVolumeSize < minVMVolume || opts.vmVolumeSize > maxVMVolume {
			return fmt.Errorf("vm-volume-size must be between %.1f and %.1f", minVMVolume, maxVMVolume)
		}
		vmVolumeSize = opts.vmVolumeSize
	}

	if cmd.Flags().Changed("vm-ram") {
		if opts.vmRam < minVMRam || opts.vmRam > maxVMRam {
			return fmt.Errorf("vm-ram must be between %d and %d", minVMRam, maxVMRam)
		}
		vmRam = opts.vmRam
	}

	// Construct API request
	body := taikuncore.UpdateQuotaCommand{
		QuotaId: &opts.projectId,
	}
	body.SetServerCpu(serverCpu)
	body.SetServerDiskSize(types.GiBToB(serverDiskSize))
	body.SetServerRam(types.GiBToB(serverRam))
	body.SetVmCpu(vmCpu)
	body.SetVmVolumeSize(vmVolumeSize)
	body.SetVmRam(types.GiBToB(vmRam))

	// Send the update request
	response, err = client.Client.ProjectQuotasAPI.ProjectquotasUpdate(context.TODO()).UpdateQuotaCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
