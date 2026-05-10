package option

import "maps"

type Option = map[OptionKey]any

type ApplyOption = func(Option)

// NewOption creates a new option. If a map is provided, it is shallow-copied.
func NewOption(x ...map[OptionKey]any) Option {
	o := Option{}
	if len(x) > 0 {
		maps.Copy(o, x[0])
	}
	return o
}

// Apply applies the options to a new option.
func Apply(opts ...ApplyOption) Option {
	o := NewOption()
	return ApplyWithDefault(o, opts...)
}

// ApplyWithDefault applies the options to a default option.
func ApplyWithDefault(d Option, opts ...ApplyOption) Option {
	o := NewOption()

	// Create a shallow copy of the default option
	maps.Copy(o, d)

	for _, apply := range opts {
		apply(o)
	}

	return o
}
