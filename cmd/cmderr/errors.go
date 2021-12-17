package cmderr

import (
	"errors"
	"fmt"
)

var NegativeLimitFlagError = errors.New("The --limit flag must be positive.")
var OutputFormatInvalidError = errors.New("Unknown output format.")
var WrongIDArgumentFormatError = errors.New("The ID argument must be a number.")

func CheckFailureError(name string) error {
	return fmt.Errorf("%s is not valid.", name)
}
