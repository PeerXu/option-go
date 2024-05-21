package option

func Setter[T any](k string) func(T) ApplyOption {
	return func(v T) ApplyOption {
		return func(o Option) {
			o[k] = v
		}
	}
}

func Getter[T any](k string) func(Option) (T, error) {
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

func New[T any](k string) (func(T) ApplyOption, func(Option) (T, error)) {
	return Setter[T](k), Getter[T](k)
}

func NewMust[T any](k string) (func(T) ApplyOption, func(Option) T) {
	getE := Getter[T](k)
	return Setter[T](k), func(o Option) T {
		v, err := getE(o)
		if err != nil {
			panic(err)
		}
		return v
	}
}
