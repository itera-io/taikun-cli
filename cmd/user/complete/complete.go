package complete

import (
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/user/list"
	"github.com/spf13/cobra"
)

func CompleteArgsWithUserID(cmd *cobra.Command) {
	cmdutils.SetArgsCompletionFunc(cmd,
		func(cmd *cobra.Command, args []string, toComplete string) []string {
			users, err := list.ListUsers(&list.ListOptions{})
			if err != nil {
				return nil
			}

			completions := make([]string, len(users))
			for i, user := range users {
				completions[i] = fmt.Sprintf(
					"%s\t%s",
					user.Id,
					user.Username,
				)
			}

			return completions
		},
	)
}
