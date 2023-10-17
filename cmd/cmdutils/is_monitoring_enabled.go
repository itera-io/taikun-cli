package cmdutils

import (
	"context"
	tk "github.com/itera-io/taikungoclient"
)

// #FIXME  - Not tested? You are better than this. I am seriously disappointed in you Radek.
func IsMonitoringEnabled(projectID int32) (isMonitoringEnabled bool, err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.ServersAPI.ServersDetails(context.TODO(), projectID).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	isMonitoringEnabled = data.Project.GetIsMonitoringEnabled()
	return
}

// #FIXME  - two files, not tested, come on Radek. You are better than this...
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
