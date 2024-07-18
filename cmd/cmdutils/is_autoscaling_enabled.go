package cmdutils

import (
	"context"
	tk "github.com/itera-io/taikungoclient"
)

func IsAutoscalingEnabled(projectID int32) (autoscalingEnabled bool, err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.ServersAPI.ServersDetails(context.TODO(), projectID).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	autoscalingEnabled = data.Project.GetIsAutoscalingEnabled()
	return
}

//import (
//	"context"
//	tk "github.com/itera-io/taikungoclient"
//)
//
//func IsAutoscalingEnabled(projectID int32) (autoscalingEnabled bool, err error) {
//	myApiClient := tk.NewClient()
//	data, response, err := myApiClient.Client.ServersAPI.ServersDetails(context.TODO(), projectID).Execute()
//	if err != nil {
//		err = tk.CreateError(response, err)
//		return
//	}
//	autoscalingEnabled = data.Project.GetIsAutoscalingEnabled()
//	return
//}

//func IsAutoscalingEnabled(projectID int32) (autoscalingEnabled bool, err error) {
//	myApiClient := tk.NewClient()
//	data, response, err := myApiClient.Client.ProjectsAPI.ProjectsList(context.TODO()).Id(projectID).Execute()
//	if err != nil {
//		err = tk.CreateError(response, err)
//		return
//	} else if data.GetTotalCount() < 1 {
//		err = fmt.Errorf("the project was not found")
//		return
//	} else {
//		result := data.GetData()[0].GetIsAutoscalingEnabled()
//		if result != false {
//			// Autoscaling is enabled
//			autoscalingEnabled = true
//		} else {
//			// Autoscaling is disabled
//			autoscalingEnabled = false
//		}
//	}
//	return
//
//	/*
//		apiClient, err := taikungoclient.NewClient()
//		if err != nil {
//			return
//		}
//
//		params := servers.NewServersDetailsParams().WithV(taikungoclient.Version)
//		params = params.WithProjectID(projectID)
//
//		response, err := apiClient.Client.Servers.ServersDetails(params, apiClient)
//		if err == nil {
//			res := response.Payload.Project.IsAutoscalingEnabled
//			if res {
//				err = errors.New("Project autoscaling already enabled.")
//			}
//		}
//		return
//	*/
//}
