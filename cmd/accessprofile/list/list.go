package list

import (
	"encoding/json"
	"fmt"

	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/itera-io/taikungoclient/models"
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
		Short: "List access profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
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
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := access_profiles.NewAccessProfilesListParams().WithV(cmdutils.ApiVersion)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		// TODO
		//cmdutils.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		// TODO
		//params = params.WithSortBy(&opts.SortBy).WithSortDirection(&cmdutils.SortDirection)
		fmt.Printf("sorting by %s\n", opts.SortBy)
	}

	var accessProfiles []*models.AccessProfilesListDto
	for {
		response, err := apiClient.Client.AccessProfiles.AccessProfilesList(params, apiClient)
		if err != nil {
			return err
		}
		accessProfiles = append(accessProfiles, response.Payload.Data...)
		usersCount := int32(len(accessProfiles))
		if opts.Limit != 0 && usersCount >= opts.Limit {
			break
		}
		if usersCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&usersCount)
	}

	if opts.Limit != 0 && int32(len(accessProfiles)) > opts.Limit {
		accessProfiles = accessProfiles[:opts.Limit]
	}

	// TODO
	//cmdutils.PrettyPrint(accessProfiles)
	jsonPayload, _ := json.MarshalIndent(accessProfiles, "", "    ")
	fmt.Println(string(jsonPayload))
	return
}
