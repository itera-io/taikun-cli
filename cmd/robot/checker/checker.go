package checker

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
}

func NewCmdCreateRobotChecker() *cobra.Command {
	var opts CreateOptions

	cmd := cobra.Command{
		Use:   "checker <NAME>",
		Short: "Dry-run for create a new robot user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkCreateRobot(args[0], &opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.AccountID, "account-id", "a", 0, "Account ID")
	cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)
	cmd.Flags().StringVarP(&opts.Description, "description", "d", "", "Robot user description")
	cmd.Flags().StringVarP(&opts.ExpiresAt, "expires-at", "e", "", "Lifetime of robot user")
	_ = cmd.MarkFlagRequired("account-id")
	_ = cmd.MarkFlagRequired("description")
	_ = cmd.MarkFlagRequired("expires-at")

	return &cmd
}

func checkCreateRobot(robotName string, opts *CreateOptions) (err error) {
	// converting time from string to time.time
	expiresAt, err := time.Parse(time.RFC3339, opts.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to parse expiration time: %w", err)
	}

	myApiClient := tk.NewClient()

	body := taikuncore.CheckerRobotCommand{
		Name:           *taikuncore.NewNullableString(&robotName),
		AccountId:      *taikuncore.NewNullableInt32(&opts.AccountID),
		OrganizationId: *taikuncore.NewNullableInt32(&opts.OrganizationID),
		Description:    *taikuncore.NewNullableString(&opts.Description),
		ExpiresAt:      *taikuncore.NewNullableTime(&expiresAt),
	}

	response, err := myApiClient.Client.RobotAPI.RobotChecker(context.TODO()).CheckerRobotCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
