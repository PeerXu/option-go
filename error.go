package option

import (
	"fmt"

	errors "github.com/PeerXu/errors-go"
)

var (
	ErrUnexpectedType = fmt.Errorf("unexpected type")
)

func ErrUnexpectedTypeFn(expect any, actual any) error {
	return fmt.Errorf("%w, expect=%T, actual=%T", ErrUnexpectedType, expect, actual)
}

var (
	ErrOptionRequired, ErrOptionRequiredFn = errors.NewErrorAndErrorFunc[string]("option required")
)
