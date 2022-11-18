package json

import (
	"encoding/json"

	"git.code.oa.com/ad/go-project-base/config/encoder"
)

type jsonEncoder struct{}

// Encode ...
func (j jsonEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode ...
func (j jsonEncoder) Decode(d []byte, v interface{}) error {
	return json.Unmarshal(d, v)
}

// String ...
func (j jsonEncoder) String() string {
	return "json"
}

// NewEncoder ...
func NewEncoder() encoder.Encoder {
	return jsonEncoder{}
}
