package whoami

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils/options"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient/client/users"
	"github.com/spf13/cobra"
)

type WhoAmIOptions struct {
	OutputOptions options.Outputter
}

func NewCmdWhoAmI(outputOptions options.Outputter) *cobra.Command {
	opts := WhoAmIOptions{
		OutputOptions: outputOptions,
	}

	cmd := cobra.Command{
		Use:   "whoami",
		Short: "Print username",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return whoAmIRun(&opts)
		},
	}

	return &cmd
}

func whoAmIRun(opts *WhoAmIOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDetailsParams().WithV(api.Version)

	response, err := apiClient.Client.Users.UsersDetails(params, apiClient)
	if err == nil {
		out.Println(opts.OutputOptions, response.Payload.Data.Username)
	}

	return
}
