package types

import (
	"github.com/itera-io/taikun-cli/utils/gmap"
)

var KubeconfigRoles = gmap.New(
	map[string]interface{}{
		"cluster-admin": int32(1),
		"admin":         int32(2),
		"edit":          int32(3),
		"view":          int32(4),
	},
)

func GetKubeconfigRole(role string) int32 {
	return KubeconfigRoles.Get(role).(int32)
}

const (
	KubeconfigAccessPersonal = "personal"
	KubeconfigAccessManagers = "managers"
	KubeconfigAccessAll      = "all"
)

var KubeconfigAccessScopes = gmap.New(
	map[string]interface{}{
		KubeconfigAccessPersonal: struct{}{},
		KubeconfigAccessManagers: struct{}{},
		KubeconfigAccessAll:      struct{}{},
	},
)
