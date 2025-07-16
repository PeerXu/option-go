package option

// Must is a helper function that returns a MustGetOptionFunc.
func Must[T any](getter GetOptionFunc[T]) MustGetOptionFunc[T] {
	return func(o Option) T {
		v, _ := getter(o)
		return v
	}
}
