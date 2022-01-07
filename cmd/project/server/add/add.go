package add

import "github.com/spf13/cobra"

type AddOptions struct {
	Count                int32
	DiskSize             int32
	Flavor               string
	KubernetesNodeLabels []string
	Name                 string
	Role                 string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a server",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return addRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// FIXME
	return
}
