package complete

import (
	"context"
	"errors"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/usertoken/list"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

func EndpointsCompleteFunc(cmd *cobra.Command, args []string, toComplete string) []string {
	myApiClient := tk.NewClient()
	limit := int32(2000)
	data, _, err := myApiClient.Client.UserTokenAPI.UsertokenAvailableEndpoints(context.TODO()).Limit(limit).Execute()
	if err != nil {
		return []string{}
	}
	completions := make([]string, 0)
	for i := 0; i < len(data.GetData()); i++ {
		res := data.GetData()[i]
		endpoint := EndpointFormatToString(res)
		completions = append(completions, endpoint)
	}

	return completions
}

// Functions for autocompletion in bin and unbind command. TOFIX
/*
func BindingEndpointsCompleteFunc(cmd *cobra.Command, args []string, toComplete string) []string {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return []string{}
	}

	TokenID, err := UserTokenIDFromUserTokenName(args[0])
	if err != nil {
		return []string{}
	}

	limit := int32(2000)
	params := user_token.NewUserTokenAvailableEndpointListParams().WithV(taikungoclient.Version).WithLimit(&limit).WithID(&TokenID)

	response, err := apiClient.Client.UserToken.UserTokenAvailableEndpointList(params, apiClient)
	if err != nil {
		return []string{}
	}

	completions := make([]string, 0)

	for i := 0; i < len(response.GetPayload().Data); i++ {
		res := response.Payload.Data[i]
		if res.ID == -1 { // case the endpoint is not bind
			endpoint := EndpointFormatToString(*res)
			completions = append(completions, endpoint)
		}
	}

	return completions
}

func UnbindingEndpointsCompleteFunc(cmd *cobra.Command, args []string, toComplete string) []string {
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
		if res.ID > 0 {
			endpoint := EndpointFormatToString(*res)
			completions = append(completions, endpoint)
		}
	}

	return completions
}
*/

func StringToEndpointFormat(endpoint string) *taikuncore.AvailableEndpointData {
	myApiClient := tk.NewClient()
	limit := int32(2000)
	data, _, err := myApiClient.Client.UserTokenAPI.UsertokenAvailableEndpoints(context.TODO()).Limit(limit).Execute()
	if err != nil {
		return nil
	}
	for i := 0; i < len(data.GetData()); i++ {
		result := data.GetData()[i]
		if endpoint == result.GetController()+"/"+result.GetMethod()+"/"+result.GetPath() {
			res := taikuncore.AvailableEndpointData{}
			res.SetController(result.GetController())
			res.SetDescription(result.GetDescription())
			res.SetId(result.GetId())
			res.SetMethod(result.GetMethod())
			res.SetPath(result.GetPath())
			return &res
		}
	}

	return nil
}

func EndpointFormatToString(res taikuncore.EndpointElements) string {
	return res.GetController() + "/" + res.GetMethod() + "/" + res.GetPath()
}

//func EndpointFormatToString(res models.EndpointElements) string {
//	return res.Controller + "/" + res.Method + "/" + res.Path
//}

func CompleteArgsWithUserTokenName(cmd *cobra.Command) {
	cmdutils.SetArgsCompletionFunc(cmd,
		func(cmd *cobra.Command, args []string, toComplete string) []string {
			users, err := list.ListUserTokens(&list.ListOptions{})
			if err != nil {
				return nil
			}

			completions := make([]string, len(users))
			for i, usertoken := range users {
				completions[i] = usertoken.GetName()
			}

			return completions
		},
	)
}

func UserTokenIDFromUserTokenName(userTokenName string) (userTokenID string, err error) {
	opts := list.ListOptions{}

	userTokenList, err := list.ListUserTokens(&opts)
	if err != nil {
		return
	}

	for i := 0; i < len(userTokenList); i++ {
		if userTokenList[i].GetName() == userTokenName {
			userTokenID = userTokenList[i].GetId()
			return
		}
	}

	err = errors.New("No user token found with name '" + userTokenName + "'.")
	return
}
