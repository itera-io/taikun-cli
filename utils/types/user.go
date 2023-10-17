package types

import (
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/utils/gmap"
	"strings"
)

//var UserRolesOld = gmap.New(
//	map[string]interface{}{
//		"user":    models.UserRole(400),
//		"manager": models.UserRole(200),
//	},
//)

// Older version, function is used only in user/add
//func GetUserRoleOld(role string) models.UserRole {
//	model, _ := UserRoles.Get(role).(models.UserRole)
//	return model
//}

func GetUserRoles() gmap.GenericMap {
	roleEnum := taikuncore.AllowedUserRoleEnumValues
	roleMap := make(map[string]interface{})
	for _, value := range roleEnum {
		roleMap[strings.ToLower(string(value))] = 0
	}
	roleGmap := gmap.New(roleMap)

	return roleGmap
}
