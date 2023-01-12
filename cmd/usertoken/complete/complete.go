package complete

import (
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/user/list"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/user_token"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func EndpointsCompleteFunc(cmd *cobra.Command, args []string, toComplete string) []string {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return []string{}
	}

	limit := int32(2000)
	params := user_token.NewUserTokenAvailableEndpointListParams().WithV(taikungoclient.Version).WithLimit(&limit)

	response, err := apiClient.Client.UserToken.UserTokenAvailableEndpointList(params, apiClient)
	if err != nil {
		return []string{}
	}

	completions := make([]string, 0)

	for i := 0; i < len(response.GetPayload().Data); i++ {
		res := response.Payload.Data[i]
		endpoint := EndpointFormatToString(*res)
		completions = append(completions, endpoint)
	}

	return completions
}

func StringToEndpointFormat(endpoint string) *models.AvailableEndpointData {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return nil
	}

	limit := int32(2000)
	params := user_token.NewUserTokenAvailableEndpointListParams().WithV(taikungoclient.Version).WithLimit(&limit)

	response, err := apiClient.Client.UserToken.UserTokenAvailableEndpointList(params, apiClient)
	if err != nil {
		return nil
	}

	for i := 0; i < len(response.GetPayload().Data); i++ {
		result := response.Payload.Data[i]
		if endpoint == result.Controller+"/"+result.Method+"/"+result.Path {
			res := models.AvailableEndpointData{}
			res.Controller = result.Controller
			res.Description = result.Description
			res.ID = result.ID
			res.Method = result.Method
			res.Path = result.Path
			return &res
		}
	}

	return nil
}

func EndpointFormatToString(res models.EndpointElements) string {
	return res.Controller + "/" + res.Method + "/" + res.Path
}

func CompleteArgsWithUserTokenName(cmd *cobra.Command) {
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
					user.ID,
					user.Username,
				)
			}

			return completions
		},
	)
}
