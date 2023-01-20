package cmderr

import (
	"errors"
	"fmt"
)

func ErrCheckFailure(name string) error {
	return fmt.Errorf("%s is not valid.", name)
}

var NoNameAutoscaler = errors.New(
	"Please specify a name for the autoscaler.",
)

var ErrNegativeLimit = errors.New(
	"The --limit flag must be positive.",
)

var ErrUnknownOutputFormat = errors.New(
	"Unknown output format.",
)

var ErrIDArgumentNotANumber = errors.New(
	"The ID argument must be a number.",
)

var ErrUnknownDateFormat = errors.New(
	"Please enter a valid date in the format dd/mm/yyyy",
)

var ErrRouterIDInvalidRange = errors.New(
	"Please specify a positive number between 1 and 255 included",
)

func ResourceNotFoundError(resourceName string, id interface{}) error {
	return fmt.Errorf("%s with ID %v not found", resourceName, id)
}

func MutuallyExclusiveFlagsError(flagA string, flagB string) error {
	return fmt.Errorf("The flags %s and %s are mutually exclusive", flagA, flagB)
}

var ErrServerHasNoFlavors = errors.New(
	"Server has no listed flavor",
)

func UnknownFlagValueError(flag string, received string, expected []string) error {
	return fmt.Errorf("unknown %s: %s, expected one of %v.", flag, received, expected)
}

var ErrProjectBackupAlreadyDisabled = errors.New(
	"Project backup already disabled",
)

var ErrProjectMonitoringAlreadyDisabled = errors.New(
	"Project monitoring already disabled",
)

var ErrProjectMonitoringAlreadyEnabled = errors.New(
	"Project monitoring already enabled",
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
