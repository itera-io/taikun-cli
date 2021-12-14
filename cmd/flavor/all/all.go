package all

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AllOptions struct {
	CloudCredentialID int32
	Limit             int32
}

func NewCmdAll() *cobra.Command {
	var opts AllOptions

	cmd := &cobra.Command{
		Use:   "all <cloud-credential-id>",
		Short: "List all flavors by cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cloudCredentialID, err := cmdutils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given ID must be a number")
			}
			opts.CloudCredentialID = cloudCredentialID
			return allRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results")

	return cmd
}

func allRun(opts *AllOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := cloud_credentials.NewCloudCredentialsAllFlavorsParams().WithV(cmdutils.ApiVersion)
	params = params.WithCloudID(opts.CloudCredentialID)

	flavors := []*models.FlavorsListDto{}
	for {
		response, err := apiClient.Client.CloudCredentials.CloudCredentialsAllFlavors(params, apiClient)
		if err != nil {
			return err
		}
		flavors = append(flavors, response.Payload.Data...)
		flavorsCount := int32(len(flavors))
		if opts.Limit != 0 && flavorsCount >= opts.Limit {
			break
		}
		if flavorsCount == response.Payload.TotalCount {
			break
		}
	}

	if opts.Limit != 0 && int32(len(flavors)) > opts.Limit {
		flavors = flavors[:opts.Limit]
	}

	cmdutils.PrettyPrint(flavors)
	return
}
