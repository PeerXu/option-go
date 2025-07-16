package option

// OptionKey is a key for option.
type OptionKey struct {
	key string
}

func (k OptionKey) String() string {
	return k.key
}
