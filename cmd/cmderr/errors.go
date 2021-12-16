package cmderr

import "errors"

var NegativeLimitFlagError = errors.New("The --limit flag must be positive.")
var OutputFormatInvalidError = errors.New("Unknown output format.")
var WrongIDArgumentFormatError = errors.New("The ID argument must be a number.")
