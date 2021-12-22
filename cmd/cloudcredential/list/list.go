package list

import (
	"fmt"

	aws "github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/list"
	azure "github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/list"
	openstack "github.com/itera-io/taikun-cli/cmd/cloudcredential/openstack/list"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID       int32
	ReverseSortDirection bool
	SortBy               string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmd.Flags().StringVarP(&opts.SortBy, "sort-by", "s", "", "Sort results by attribute value")
	cmdutils.AddLimitFlag(cmd)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	err = openstack.ListRun((*openstack.ListOptions)(opts))
	if err != nil {
		return
	}
	fmt.Println()
	err = azure.ListRun((*azure.ListOptions)(opts))
	if err != nil {
		return
	}
	fmt.Println()
	err = aws.ListRun((*aws.ListOptions)(opts))
	return
}
