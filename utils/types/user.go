package types

import "github.com/itera-io/taikungoclient/models"

var UserRoles = map[string]interface{}{
	"user":    400,
	"manager": 200,
}

func GetUserRole(role string) models.UserRole {
	return UserRoles[role].(models.UserRole)
}
