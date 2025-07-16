package option

// LightNew returns a setter and getter.
func LightNew[T any](ks ...string) (SetOptionFunc[T], GetOptionFunc[T]) {
	_, setter, getter := new[T](ks...)
	return setter, getter
}

// LightMustNew returns a setter and must getter.
func LightMustNew[T any](ks ...string) (SetOptionFunc[T], MustGetOptionFunc[T]) {
	_, setter, getter := new[T](ks...)
	return setter, Must(getter)
}

// New returns a key, setter and getter.
func New[T any](ks ...string) (OptionKey, SetOptionFunc[T], GetOptionFunc[T]) {
	return new[T](ks...)
}

// NewMust returns a key, setter and must getter.
func NewMust[T any](ks ...string) (OptionKey, SetOptionFunc[T], MustGetOptionFunc[T]) {
	ok, setter, getter := new[T](ks...)
	return ok, setter, Must(getter)
}

// new is a low level API, it returns the key, setter and getter.
func new[T any](ks ...string) (OptionKey, SetOptionFunc[T], GetOptionFunc[T]) {
	var k string
	if len(ks) > 0 {
		k = ks[0]
	}
	sk := OptionKey{key: k}
	return sk, Setter[T](sk), Getter[T](sk)
}
