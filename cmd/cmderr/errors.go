package cmderr

import (
	"errors"
	"fmt"
)

func ErrCheckFailure(name string) error {
	return fmt.Errorf("%s is not valid", name)
}

var ErrNoNameAutoscaler = errors.New(
	"please specify a name for the autoscaler",
)

var ErrNegativeLimit = errors.New(
	"the --limit flag must be positive",
)

var ErrUnknownOutputFormat = errors.New(
	"unknown output format",
)

var ErrIDArgumentNotANumber = errors.New(
	"the ID argument must be a number",
)

var ErrUnknownDateFormat = errors.New(
	"please enter a valid date in the format dd.mm.yyyy, dd.mm.yyyy hh:mm, or dd.mm.yyyy hh:mm:ss",
)

var ErrRouterIDInvalidRange = errors.New(
	"please specify a positive number between 1 and 255 included",
)

func ResourceNotFoundError(resourceName string, id interface{}) error {
	return fmt.Errorf("%s with ID %v not found", resourceName, id)
}

func MutuallyExclusiveFlagsError(flagA string, flagB string) error {
	return fmt.Errorf("the flags %s and %s are mutually exclusive", flagA, flagB)
}

func UnknownFlagValueError(flag string, received string, expected []string) error {
	return fmt.Errorf("unknown %s: %s, expected one of %v", flag, received, expected)
}

var ErrProjectBackupAlreadyDisabled = errors.New(
	"project backup already disabled",
)

var ErrProjectMonitoringAlreadyDisabled = errors.New(
	"project monitoring already disabled",
)

var ErrProjectMonitoringAlreadyEnabled = errors.New(
	"project monitoring already enabled",
)

func ProgramError(functionName string, err error) error {
	return fmt.Errorf(
		"%s: %s\n"+
			"This is a bug with the CLI. "+
			"If the issue persists, please report it at "+
			"https://github.com/itera-io/taikun-cli/issues",
		functionName,
		err,
	)
}
