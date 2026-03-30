package regenerate

import (
	"context"
	"fmt"
	"time"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type RegenerateOptions struct {
	ExpiresAt string
}

func NewCmdRobotRegenerateTokens() *cobra.Command {
	var opts RegenerateOptions

	cmd := cobra.Command{
		Use:   "regenerate <ROBOT_ID>",
		Short: "Regenerate tokens for a robot user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return regenerateRobotTokens(args[0], &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.ExpiresAt, "expires-at", "e", "", "Lifetime of robot user")
	_ = cmd.MarkFlagRequired("expires-at")

	return &cmd
}

func regenerateRobotTokens(robotID string, opts *RegenerateOptions) (err error) {
	// converting time from string to time.time
	expiresAt, err := time.Parse(time.RFC3339, opts.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to parse expiration time: %w", err)
	}

	myApiClient := tk.NewClient()

	body := taikuncore.RegenerateRobotTokenCommand{
		Id:        *taikuncore.NewNullableString(&robotID),
		ExpiresAt: *taikuncore.NewNullableTime(&expiresAt),
	}

	data, response, err := myApiClient.Client.RobotAPI.RobotRegenerate(context.TODO()).RegenerateRobotTokenCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	fmt.Printf("{\n\tACCESS_KEY: %s,\n\tSECRET_KEY: %s,\n}\n", *data.AccessKey.Get(), *data.SecretKey.Get())
	return nil
}
