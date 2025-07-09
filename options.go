package option

import "github.com/jinzhu/copier"

type Option = map[structKey]any

type ApplyOption = func(Option)

func NewOption(x ...map[structKey]any) Option {
	if len(x) > 0 {
		return x[0]
	}
	return map[structKey]any{}
}

func Apply(opts ...ApplyOption) Option {
	o := NewOption()
	return ApplyWithDefault(o, opts...)
}

func ApplyWithDefault(d Option, opts ...ApplyOption) Option {
	o := NewOption()
	copier.CopyWithOption(&o, &d, copier.Option{DeepCopy: true})
	for _, apply := range opts {
		apply(o)
	}
	return o
}
