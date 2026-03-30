package create

import (
	"context"
	"fmt"
	"time"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	AccountID      int32
	OrganizationID int32
	Description    string
	ExpiresAt      string
	Scopes         []string
	IPs            []string
}

func NewCmdCreateRobot() *cobra.Command {
	opts := CreateOptions{
		Scopes: make([]string, 0),
		IPs:    make([]string, 0),
	}

	cmd := cobra.Command{
		Use:   "create <NAME>",
		Short: "Create a new robot user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createRobot(args[0], &opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.AccountID, "account-id", "a", 0, "Account ID")
	cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)
	cmd.Flags().StringVarP(&opts.Description, "description", "d", "", "Robot user description")
	cmd.Flags().StringVarP(&opts.ExpiresAt, "expires-at", "e", "", "Lifetime of robot user")
	cmd.Flags().StringSliceVarP(&opts.Scopes, "scope", "s", nil, "Scope of the robot user")
	cmd.Flags().StringSliceVarP(&opts.IPs, "ip", "i", nil, "IPs of the robot user")
	_ = cmd.MarkFlagRequired("account-id")
	_ = cmd.MarkFlagRequired("description")
	_ = cmd.MarkFlagRequired("expires-at")

	return &cmd
}

func createRobot(robotName string, opts *CreateOptions) (err error) {
	// converting time from string to time.time
	expiresAt, err := time.Parse(time.RFC3339, opts.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to parse expiration time: %w", err)
	}

	myApiClient := tk.NewClient()

	body := taikuncore.CreateRobotUserCommand{
		Name:           *taikuncore.NewNullableString(&robotName),
		AccountId:      *taikuncore.NewNullableInt32(&opts.AccountID),
		OrganizationId: *taikuncore.NewNullableInt32(&opts.OrganizationID),
		Description:    *taikuncore.NewNullableString(&opts.Description),
		ExpiresAt:      *taikuncore.NewNullableTime(&expiresAt),
		Scopes:         opts.Scopes,
		Ips:            opts.IPs,
	}

	data, response, err := myApiClient.Client.RobotAPI.RobotCreate(context.TODO()).CreateRobotUserCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	fmt.Printf("{\n\tACCESS_KEY: %s,\n\tSECRET_KEY: %s,\n}\n", *data.AccessKey.Get(), *data.SecretKey.Get())
	return nil
}
