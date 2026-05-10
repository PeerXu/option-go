package option

// Must is a helper function that returns a MustGetOptionFunc that panics on error.
func Must[T any](getter GetOptionFunc[T]) MustGetOptionFunc[T] {
	return func(o Option) T {
		v, err := getter(o)
		if err != nil {
			panic(err)
		}
		return v
	}
}
