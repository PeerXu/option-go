package option

import "github.com/jinzhu/copier"

type Option = map[OptionKey]any

type ApplyOption = func(Option)

// NewOption creates a new option.
func NewOption(x ...map[OptionKey]any) Option {
	if len(x) > 0 {
		return x[0]
	}
	return map[OptionKey]any{}
}

// Apply applies the options to a new option.
func Apply(opts ...ApplyOption) Option {
	o := NewOption()
	return ApplyWithDefault(o, opts...)
}

// ApplyWithDefault applies the options to a default option.
func ApplyWithDefault(d Option, opts ...ApplyOption) Option {
	o := NewOption()
	copier.CopyWithOption(&o, &d, copier.Option{DeepCopy: true})
	for _, apply := range opts {
		apply(o)
	}
	return o
}
