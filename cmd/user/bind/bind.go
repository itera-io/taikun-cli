package bind

import (
	"fmt"
	"taikun-cli/api"

	"github.com/spf13/cobra"
)

type BindOptions struct {
	Username  string
	ProjectID int
}

func NewCmdBind(apiClient *api.Client) *cobra.Command {
	var opts BindOptions

	cmd := &cobra.Command{
		Use:   "bind",
		Short: "Bind a user to a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return bindRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "Username (required)")
	cmd.MarkFlagRequired("username")

	cmd.Flags().IntVarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmd.MarkFlagRequired("project-id")

	return cmd
}

func bindRun(opts *BindOptions) (err error) {
	fmt.Printf("Binding user %s to project with ID %d\n", opts.Username, opts.ProjectID)
	// FIXME
	return
}
