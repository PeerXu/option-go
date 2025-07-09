package option

type structKey struct {
	key string
}

func (k structKey) String() string {
	return k.key
}
