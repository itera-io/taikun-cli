package list

import (
	"fmt"

	aws "taikun-cli/cmd/cloudcredential/aws/list"
	azure "taikun-cli/cmd/cloudcredential/azure/list"
	openstack "taikun-cli/cmd/cloudcredential/openstack/list"
	"taikun-cli/config"
	"taikun-cli/utils"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit                int32
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
			if opts.Limit < 0 {
				return utils.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return config.OutputFormatInvalidError
			}
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmd.Flags().StringVarP(&opts.SortBy, "sort-by", "s", "", "Sort results by attribute value")

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
