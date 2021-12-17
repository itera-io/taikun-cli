package types

import "github.com/itera-io/taikungoclient/models"

var UserRoles = map[string]interface{}{
	"user":    models.UserRole(400),
	"manager": models.UserRole(200),
}

func GetUserRole(role string) models.UserRole {
	return UserRoles[role].(models.UserRole)
}
