package option

type GetOptionFunc[T any] func(Option) (T, error)
type MustGetOptionFunc[T any] func(Option) T
type SetOptionFunc[T any] func(T) ApplyOption

// Setter is a helper function that returns a SetOptionFunc.
func Setter[T any](k OptionKey) SetOptionFunc[T] {
	return func(v T) ApplyOption {
		return func(o Option) {
			o[k] = v
		}
	}
}

// Getter is a helper function that returns a GetOptionFunc.
func Getter[T any](k OptionKey) GetOptionFunc[T] {
	return func(o Option) (T, error) {
		var x T
		i := o[k]
		if i == nil {
			return x, ErrOptionRequiredFn(k)
		}

		v, ok := i.(T)
		if !ok {
			return x, ErrUnexpectedTypeFn(x, i)
		}

		return v, nil
	}
}
