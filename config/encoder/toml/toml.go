package toml

import (
	"bytes"

	"git.code.oa.com/ad/go-project-base/config/encoder"
	"github.com/BurntSushi/toml"
)

type tomlEncoder struct{}

// Encode ...
func (t tomlEncoder) Encode(v interface{}) ([]byte, error) {
	b := bytes.NewBuffer(nil)
	defer b.Reset()
	err := toml.NewEncoder(b).Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Decode ...
func (t tomlEncoder) Decode(d []byte, v interface{}) error {
	return toml.Unmarshal(d, v)
}

// String ...
func (t tomlEncoder) String() string {
	return "toml"
}

// NewEncoder ...
func NewEncoder() encoder.Encoder {
	return tomlEncoder{}
}
