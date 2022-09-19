package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
	"github.com/itera-io/taikungoclient/models"
)

var UserRoles = gmap.New(
	map[string]interface{}{
		"user":    models.UserRole(400),
		"manager": models.UserRole(200),
	},
)

func GetUserRole(role string) models.UserRole {
	model, _ := UserRoles.Get(role).(models.UserRole)
	return model
}
