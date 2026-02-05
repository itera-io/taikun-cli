package notifications

import (
	"context"
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type NotificationsOptions struct {
	Mode string
}

func NewCmdNotifications() *cobra.Command {
	var opts NotificationsOptions

	cmd := cobra.Command{
		Use:   "notifications [enable|disable]",
		Short: "Toggle notification mode",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				if opts.Mode != "" {
					return fmt.Errorf("provide mode either as argument or --mode, not both")
				}
				opts.Mode = args[0]
			}
			return notificationsRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Mode, "mode", "m", "", "Notification mode (enable|disable)")
	cmdutils.SetFlagCompletionValues(&cmd, "mode", "enable", "disable")

	return &cmd
}

func notificationsRun(opts *NotificationsOptions) (err error) {
	mode, err := normalizeMode(opts.Mode)
	if err != nil {
		return err
	}

	myApiClient := tk.NewClient()
	body := taikuncore.ToggleNotificationModeCommand{}
	body.SetMode(mode)

	response, err := myApiClient.Client.UsersAPI.UsersToggleNotificationMode(context.TODO()).
		ToggleNotificationModeCommand(body).
		Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return
}

func normalizeMode(mode string) (string, error) {
	if mode == "" {
		return "", fmt.Errorf("mode is required (enable|disable)")
	}

	switch strings.ToLower(mode) {
	case "enable", "enabled", "on", "true", "1":
		return "enable", nil
	case "disable", "disabled", "off", "false", "0":
		return "disable", nil
	default:
		return "", fmt.Errorf("invalid mode %q (enable|disable)", mode)
	}
}
