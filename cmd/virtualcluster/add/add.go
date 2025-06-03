package add

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

type AddOptions struct {
	parentProjectId    int32
	virtualClusterName string
	alertingProfileId  int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <parent-id> <name>",
		Short: "Create a new virtual cluster",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.parentProjectId, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.virtualClusterName = args[1]

			return addRun(&opts)
		},
		Args: cobra.ExactArgs(2),
	}

	cmd.Flags().Int32VarP(&opts.alertingProfileId, "alerting-profile-id", "a", 0, "Alerting profile ID.")

	return &cmd
}

// addRun calls the API with a custom body from arguments. It than prints the result.
func addRun(opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()

	if opts.alertingProfileId == 0 {
		// Get alerting profile ID of parent
		data, response, err := myApiClient.Client.ProjectsAPI.ProjectsList(context.TODO()).Id(opts.parentProjectId).Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		if data.GetTotalCount() == 1 {
			opts.alertingProfileId = data.GetData()[0].GetAlertingProfileId()
		} else {
			return fmt.Errorf("cannot get parent project alerting profile id")
		}

	}

	body := taikuncore.CreateVirtualClusterCommand{
		ProjectId:         &opts.parentProjectId,
		Name:              &opts.virtualClusterName,
		AlertingProfileId: *taikuncore.NewNullableInt32(&opts.alertingProfileId),
	}

	response, err := myApiClient.Client.VirtualClusterAPI.VirtualClusterCreate(context.TODO()).CreateVirtualClusterCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
