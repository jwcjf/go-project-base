package yaml

import (
	"git.code.oa.com/ad/go-project-base/config/encoder"
	"github.com/ghodss/yaml"
)

type yamlEncoder struct{}

// Encode ...
func (y yamlEncoder) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// Decode ...
func (y yamlEncoder) Decode(d []byte, v interface{}) error {
	return yaml.Unmarshal(d, v)
}

// String ...
func (y yamlEncoder) String() string {
	return "yaml"
}

// NewEncoder ...
func NewEncoder() encoder.Encoder {
	return yamlEncoder{}
}
