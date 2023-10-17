package types

import (
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/utils/gmap"
)

//var ServerRoles = gmap.New(
//	map[string]interface{}{
//		"bastion":    models.CloudRole(100),
//		"kubemaster": models.CloudRole(200),
//		"kubeworker": models.CloudRole(300),
//	},
//)

//func GetServerRole(serverRole string) models.CloudRole {
//	model, _ := ServerRoles.Get(serverRole).(models.CloudRole)
//	return model
//}

var ServerRoles = gmap.New(
	map[string]interface{}{
		"bastion":    taikuncore.CLOUDROLE_BASTION,
		"kubemaster": taikuncore.CLOUDROLE_KUBEMASTER,
		"kubeworker": taikuncore.CLOUDROLE_KUBEWORKER,
	},
)

func GetServerRole(serverRole string) taikuncore.CloudRole {
	model, _ := ServerRoles.Get(serverRole).(taikuncore.CloudRole)
	return model
}
