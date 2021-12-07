package noop

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdNoop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "noop",
		Short: "Do nothing",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("noop")
		},
	}
	return cmd
}
