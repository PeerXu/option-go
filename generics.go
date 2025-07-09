package option

func Setter[T any](k structKey) func(T) ApplyOption {
	return func(v T) ApplyOption {
		return func(o Option) {
			o[k] = v
		}
	}
}

func Getter[T any](k structKey) func(Option) (T, error) {
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

func New[T any](k string) (structKey, func(T) ApplyOption, func(Option) (T, error)) {
	sk := structKey{key: k}
	return sk, Setter[T](sk), Getter[T](sk)
}

func Must[T any](getter func(Option) (T, error)) func(Option) T {
	return func(o Option) T {
		v, _ := getter(o)
		return v
	}
}

func NewMust[T any](k string) (structKey, func(T) ApplyOption, func(Option) T) {
	sk := structKey{key: k}
	getE := Getter[T](sk)
	return sk, Setter[T](sk), Must(getE)
}
