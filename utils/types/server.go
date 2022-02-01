package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	"github.com/itera-io/taikungoclient/models"
)

var ServerRoles = gmap.New(
	map[string]interface{}{
		"bastion":    models.CloudRole(100),
		"kubemaster": models.CloudRole(200),
		"kubeworker": models.CloudRole(300),
	},
)

func GetServerRole(serverRole string) models.CloudRole {
	return ServerRoles.Get(serverRole).(models.CloudRole)
}
