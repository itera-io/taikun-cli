package expiration

import (
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ExtendLifetimeOptions struct {
	ProjectID          int32
	ExpirationDate     string
	DeleteOnExpiration bool
}

func NewCmdExpiration() *cobra.Command {
	var opts ExtendLifetimeOptions

	cmd := cobra.Command{
		Use:   "expiration <project-id>",
		Short: "Manage expiration for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			if opts.ExpirationDate != "" {
				if !types.StrIsValidDate(opts.ExpirationDate) {
					return cmderr.ErrUnknownDateFormat
				}
			}
			return extendProjectLifetime(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.DeleteOnExpiration, "delete-on-expiration", "del", false, "Delete on expiration (required)")
	cmdutils.MarkFlagRequired(&cmd, "delete-on-expiration")
	cmd.Flags().StringVarP(&opts.ExpirationDate, "expiration-date", "e", "", fmt.Sprintf("Expiration date in the format: %s", types.ExpectedDateFormat))

	return &cmd
}

func extendProjectLifetime(opts *ExtendLifetimeOptions) (err error) {

	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}
	body := models.ProjectExtendLifeTimeCommand{ProjectID: opts.ProjectID, DeleteOnExpiration: opts.DeleteOnExpiration}

	if opts.ExpirationDate != "" {
		expiredAt := types.StrToDateTime(opts.ExpirationDate)
		body.ExpireAt = &expiredAt
	}

	params := projects.NewProjectsExtendLifeTimeParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Projects.ProjectsExtendLifeTime(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
