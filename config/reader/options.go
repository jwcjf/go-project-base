package reader

import (
	"github.com/jwcjf/go-project-base/config/encoder"
	"github.com/jwcjf/go-project-base/config/encoder/json"
	"github.com/jwcjf/go-project-base/config/encoder/toml"
	"github.com/jwcjf/go-project-base/config/encoder/xml"
	"github.com/jwcjf/go-project-base/config/encoder/yaml"
)

// Options ...
type Options struct {
	Encoding map[string]encoder.Encoder
}

// Option ...
type Option func(o *Options)

// NewOptions ...
func NewOptions(opts ...Option) Options {
	options := Options{
		Encoding: map[string]encoder.Encoder{
			"json": json.NewEncoder(),
			"yaml": yaml.NewEncoder(),
			"toml": toml.NewEncoder(),
			"xml":  xml.NewEncoder(),
			"yml":  yaml.NewEncoder(),
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

// WithEncoder ...
func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		if o.Encoding == nil {
			o.Encoding = make(map[string]encoder.Encoder)
		}
		o.Encoding[e.String()] = e
	}
}
