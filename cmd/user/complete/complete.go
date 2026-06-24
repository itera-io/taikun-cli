package complete

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func CompleteArgsWithUserID(cmd *cobra.Command) {
	cmdutils.SetArgsCompletionFunc(cmd,
		func(cmd *cobra.Command, args []string, toComplete string) []string {
			ctx, cancel := cmdutils.APIContext(cmd)
			defer cancel()
			myApiClient := tk.NewClient()
			data, _, err := myApiClient.Client.UsersAPI.UsersDropdown(ctx).Execute()
			if err != nil {
				return nil
			}

			completions := make([]string, len(data.GetData()))
			for i, user := range data.GetData() {
				completions[i] = user.GetId()
			}

			return completions
		},
	)
}
