package whoami

import (
	"fmt"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdWhoami() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Print username",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return whoamiRun(cmd)
		},
	}
	return cmd
}

func whoamiRun(cmd *cobra.Command) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, _, err := myApiClient.Client.UsersAPI.UsersUserInfo(ctx).Execute()
	if err != nil {
		return err
	}

	// Manipulate the gathered data
	username := data.Data.GetUsername()
	fmt.Printf("%s\n", username)
	return
}
