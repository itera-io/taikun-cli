package cmderr

import (
	"errors"
	"fmt"
)

func CheckFailureError(name string) error {
	return fmt.Errorf("%s is not valid.", name)
}

var NegativeLimitFlagError = errors.New(
	"The --limit flag must be positive.",
)

var OutputFormatInvalidError = errors.New(
	"Unknown output format.",
)

var WrongIDArgumentFormatError = errors.New(
	"The ID argument must be a number.",
)

var InvalidDateFormatError = errors.New(
	"Please enter a valid date in the format dd/mm/yyyy",
)
