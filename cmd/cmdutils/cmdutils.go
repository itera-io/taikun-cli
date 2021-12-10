package cmdutils

import (
	"log"

	"github.com/spf13/cobra"
)

const ApiVersion = "1"

func MarkFlagRequired(cmd *cobra.Command, flag string) {
	if err := cmd.MarkFlagRequired(flag); err != nil {
		log.Fatal(err)
	}
}
