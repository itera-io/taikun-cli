package check

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type CheckOptions struct {
	Name string
}

func NewCmdCheckDuplicateEntity() *cobra.Command {
	var opts CheckOptions

	cmd := cobra.Command{
		Use:   "check-duplicate-entity <ACCOUNT_ID>",
		Short: "Checks if entity is duplicated",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			accountID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return checkDuplicateEntity(accountID, &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Group name")
	_ = cmd.MarkFlagRequired("name")

	return &cmd
}

func checkDuplicateEntity(accountID int32, opts *CheckOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.CheckDuplicateGroupCommand{
		AccountId: *taikuncore.NewNullableInt32(&accountID),
		Name:      *taikuncore.NewNullableString(&opts.Name),
	}

	response, err := myApiClient.Client.GroupsAPI.GroupsCheckDuplicateEntity(context.TODO()).CheckDuplicateGroupCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
