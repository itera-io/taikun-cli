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

var IDArgumentNotANumberError = errors.New(
	"The ID argument must be a number.",
)

var InvalidDateFormatError = errors.New(
	"Please enter a valid date in the format dd/mm/yyyy",
)

var RouterIDRangeError = errors.New(
	"Please specify a positive number between 1 and 255 included",
)

func ResourceNotFoundError(resourceName string, id int32) error {
	return fmt.Errorf("%s with ID %d not found", resourceName, id)
}

var ServerHasNoFlavorError = errors.New(
	"Server has no listed flavor",
)
