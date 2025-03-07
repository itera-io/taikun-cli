package enable

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type EnableOptions struct {
	ProjectID            int32
	AutoscalingGroupName string
	Flavor               string
	DiskSize             int32
	MaxSize              int32
	MinSize              int32
	Spot                 bool
}

func NewCmdEnable() *cobra.Command {
	var opts EnableOptions

	cmd := cobra.Command{
		Use:   "enable <project-id>",
		Short: "Enable autoscaling for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return enableRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AutoscalingGroupName, "autoscaler-name", "n", "", "The autoscaler's name (required)")
	cmdutils.MarkFlagRequired(&cmd, "autoscaler-name")

	cmd.Flags().StringVarP(&opts.Flavor, "autoscaler-flavor", "f", "", "The autoscaler's flavor (required)")
	cmdutils.MarkFlagRequired(&cmd, "autoscaler-flavor")
	cmdutils.SetFlagCompletionFunc(&cmd, "autoscaler-flavor", cmdutils.FlavorCompletionFunc)

	cmd.Flags().Int32Var(&opts.MaxSize, "max-size", 1, "The autoscaler's maximum size")
	cmd.Flags().Int32Var(&opts.MinSize, "min-size", 1, "The autoscaler's minimum size")
	cmd.Flags().Int32Var(&opts.DiskSize, "disk-size", 30, "The autoscaler's disk size")
	cmd.Flags().BoolVarP(&opts.Spot, "spot-enable", "s", false, "Use to enable spot flavors for autoscaler (default false)")

	return &cmd
}

func enableRun(opts *EnableOptions) (err error) {
	autoscalingEnabled, err := cmdutils.IsAutoscalingEnabled(opts.ProjectID)
	if err != nil {
		return err
	}
	if autoscalingEnabled {
		err = fmt.Errorf("Project autoscaling already enabled")
		return err
	}

	myApiClient := tk.NewClient()
	diskSize := types.GiBToB(opts.DiskSize)
	body := taikuncore.EnableAutoscalingCommand{
		Id:                   &opts.ProjectID,
		AutoscalingGroupName: *taikuncore.NewNullableString(&opts.AutoscalingGroupName),
		MinSize:              &opts.MinSize,
		MaxSize:              &opts.MaxSize,
		DiskSize:             &diskSize,
		Flavor:               *taikuncore.NewNullableString(&opts.Flavor),
		SpotEnabled:          &opts.Spot,
	}
	response, err := myApiClient.Client.AutoscalingAPI.AutoscalingEnable(context.TODO()).EnableAutoscalingCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
