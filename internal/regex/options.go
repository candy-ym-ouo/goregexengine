package regex

import "strings"

// Options controls compilation and execution.
type Options struct {
	IgnoreCase, Multiline, DotAll, Strict bool
	ForceBacktracking                     bool
}

// ParseOptions converts the compact flags string into options.
func ParseOptions(flags string) Options {
	var o Options
	for _, f := range strings.ToLower(flags) {
		switch f {
		case 'i':
			o.IgnoreCase = true
		case 'm':
			o.Multiline = true
		case 's':
			o.DotAll = true
		case 'x':
			o.Strict = true
		}
	}
	if strings.Contains(flags, "strict") {
		o.Strict = true
	}
	return o
}

func (o Options) regexpFlags() string {
	var b strings.Builder
	if o.IgnoreCase {
		b.WriteByte('i')
	}
	if o.Multiline {
		b.WriteByte('m')
	}
	if o.DotAll {
		b.WriteByte('s')
	}
	return b.String()
}
