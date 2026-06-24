package regenerate

import (
	"fmt"
	"time"

	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
			return regenerateRobotTokens(cmd, args[0], &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.ExpiresAt, "expires-at", "e", "", "Lifetime of robot user")
	_ = cmd.MarkFlagRequired("expires-at")

	return &cmd
}

func regenerateRobotTokens(cmd *cobra.Command, robotID string, opts *RegenerateOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

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

	data, response, err := myApiClient.Client.RobotAPI.RobotRegenerate(ctx).RegenerateRobotTokenCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	fmt.Printf("{\n\tACCESS_KEY: %s,\n\tSECRET_KEY: %s,\n}\n", *data.AccessKey.Get(), *data.SecretKey.Get())
	return nil
}
