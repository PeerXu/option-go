package option

import (
	"fmt"

	errors "github.com/PeerXu/errors-go/v2"
)

var (
	ErrUnexpectedType, ErrUnexpectedTypeFn = errors.NewErrorAndErrorFunc("unexpected type", func(err error, expect any, args ...any) error {
		return fmt.Errorf("%w: expect=%T, actual=%T", err, expect, args[0])
	})
	ErrOptionRequired, ErrOptionRequiredFn = errors.NewErrorAndErrorFunc[OptionKey]("option required")
)
