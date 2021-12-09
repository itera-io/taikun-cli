package noop

import (
	"fmt"
	"taikun-cli/api"

	"github.com/spf13/cobra"
)

func NewCmdNoop(apiClient *api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "noop",
		Short: "Do nothing",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("noop")
		},
	}
	return cmd
}
