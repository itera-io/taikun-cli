package expiration

import (
	"fmt"
	"time"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type ExtendLifetimeOptions struct {
	ProjectID          int32
	ExpirationDate     string
	DeleteOnExpiration bool
	RemoveExpiration   bool
}

func NewCmdExpiration() *cobra.Command {
	var opts ExtendLifetimeOptions

	cmd := cobra.Command{
		Use:   "expiration <project-id>",
		Short: "Manage expiration for a project. Projects can expire every full hour. Minutes and seconds are ignored. If no minutes are specified - midnight is used.",
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
			return extendProjectLifetime(cmd, &opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.RemoveExpiration, "remove-expiration", "r", false, "Clear expiration date - project never expires.")
	cmd.Flags().BoolVarP(&opts.DeleteOnExpiration, "delete-on-expiration", "d", false, "Delete project on expiration")
	cmd.Flags().StringVarP(&opts.ExpirationDate, "expiration-date", "e", "", fmt.Sprintf("Expiration date in the format: %s, %s, or %s.  Minutes and seconds are ignored. Projects can expire only at 00 (full hour).", types.ExpectedDateFormat, "dd.mm.yyyy hh:mm", types.ExpectedDateTimeFormat))

	return &cmd
}

func extendProjectLifetime(cmd *cobra.Command, opts *ExtendLifetimeOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	body := taikuncore.ProjectExtendLifeTimeCommand{}
	body.SetProjectId(opts.ProjectID)

	if (opts.RemoveExpiration && opts.ExpirationDate != "") || (!opts.RemoveExpiration && opts.ExpirationDate == "") {
		return fmt.Errorf("specify one --remove-expiration (-r) or --expiration-date (-e). Flags mutually exclusive")
	}

	if opts.RemoveExpiration { // Remove expiration
		body.SetDeleteOnExpiration(opts.DeleteOnExpiration)
	} else { // Set expiration
		expiredAt := types.StrToDateTime(opts.ExpirationDate)
		body.SetExpireAt(time.Time(expiredAt))
		body.SetDeleteOnExpiration(opts.DeleteOnExpiration)
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.ProjectsAPI.ProjectsExtendLifetime(ctx).ProjectExtendLifeTimeCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	out.PrintStandardSuccess()
	return

}
