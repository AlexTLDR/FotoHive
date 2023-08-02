package errors

import "errors"

// Wrapping these functions in a package allows us to use them in other packages

var (
	As = errors.As
	Is = errors.Is
)
