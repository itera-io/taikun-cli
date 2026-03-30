package update

import (
	"context"
	"fmt"
	"time"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type UpdateOptions struct {
	Description string
	ExpiresAt   string
	Name        string
	IPs         []string
}

func NewCmdUpdate() *cobra.Command {
	var opts UpdateOptions

	cmd := cobra.Command{
		Use:   "update <ROBOT_ID>",
		Short: "Update robot user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return updateRobot(args[0], &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Description, "description", "d", "", "Description")
	cmd.Flags().StringVarP(&opts.ExpiresAt, "expires-at", "e", "", "Expires at (RFC3339)")
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name")
	cmd.Flags().StringSliceVarP(&opts.IPs, "ips", "i", nil, "IPs")

	return &cmd
}

func updateRobot(robotID string, opts *UpdateOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.EditRobotUserCommand{
		Id: *taikuncore.NewNullableString(&robotID),
	}

	if opts.Description != "" {
		body.SetDescription(opts.Description)
	}
	if opts.Name != "" {
		body.SetName(opts.Name)
	}
	if opts.IPs != nil {
		body.SetIps(opts.IPs)
	}
	if opts.ExpiresAt != "" {
		expiresAt, err := time.Parse(time.RFC3339, opts.ExpiresAt)
		if err != nil {
			return fmt.Errorf("failed to parse expiration time: %w", err)
		}
		body.SetExpiresAt(expiresAt)
	}

	response, err := myApiClient.Client.RobotAPI.RobotUpdate(context.TODO()).EditRobotUserCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
