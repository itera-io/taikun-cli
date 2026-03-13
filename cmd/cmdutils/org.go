package cmdutils

import (
	"fmt"
	"os"
	"strconv"
)

const EnvOrgID = "TAIKUN_ORGANIZATION_ID"

func ResolveOrgID(flagValue int32, isRobot bool) (int32, error) {
	if flagValue != 0 {
		return flagValue, nil
	}

	envID, envSet, err := readOrgIDFromEnv()
	if err != nil {
		return 0, err
	}
	if envSet {
		return envID, nil
	}

	if isRobot {
		return 0, nil
	}

	return 0, ErrMissingOrg()
}

func readOrgIDFromEnv() (int32, bool, error) {
	raw := os.Getenv(EnvOrgID)
	if raw == "" {
		return 0, false, nil
	}

	v, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		return 0, true, fmt.Errorf(
			"invalid %s value %q: must be a positive integer",
			EnvOrgID, raw,
		)
	}

	if v <= 0 {
		return 0, true, fmt.Errorf(
			"invalid %s value %q: must be a positive integer",
			EnvOrgID, raw,
		)
	}

	return int32(v), true, nil
}

func ErrMissingOrg() error {
	return fmt.Errorf(
		"organization ID is required; either pass --organization-id or set the %s environment variable",
		EnvOrgID,
	)
}
