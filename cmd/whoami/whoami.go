package whoami

import (
	"context"
	"fmt"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdWhoami() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Print username",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return whoamiRun()
		},
	}
	return cmd
}

func whoamiRun() (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, _, err := myApiClient.Client.UsersAPI.UsersUserInfo(context.TODO()).Execute()
	if err != nil {
		return err
	}

	// Manipulate the gathered data
	username := data.Data.GetUsername()
	fmt.Printf("%s\n", username)
	return
}
