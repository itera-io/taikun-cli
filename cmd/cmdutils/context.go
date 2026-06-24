package cmdutils

import (
	"context"

	"github.com/itera-io/taikun-cli/config"
	"github.com/spf13/cobra"
)

func APIContext(cmd *cobra.Command) (context.Context, context.CancelFunc) {
	return context.WithTimeout(cmd.Context(), config.APITimeout)
}
