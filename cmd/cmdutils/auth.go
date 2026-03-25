package cmdutils

import (
	"os"

	taikuncore "github.com/itera-io/taikungoclient"
)

// IsRobotAuth returns true if the current CLI session uses robot authentication.
// Robot auth is detected by the presence of both TAIKUN_ACCESS_KEY and TAIKUN_SECRET_KEY env vars.
func IsRobotAuth() bool {
	return os.Getenv(taikuncore.TaikunAccessKey) != "" && os.Getenv(taikuncore.TaikunSecretKey) != ""
}
