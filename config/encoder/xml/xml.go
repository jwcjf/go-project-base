package xml

import (
	"encoding/xml"

	"git.code.oa.com/ad/go-project-base/config/encoder"
)

type xmlEncoder struct{}

// Encode ...
func (x xmlEncoder) Encode(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

// Decode ...
func (x xmlEncoder) Decode(d []byte, v interface{}) error {
	return xml.Unmarshal(d, v)
}

// String ...
func (x xmlEncoder) String() string {
	return "xml"
}

// NewEncoder ...
func NewEncoder() encoder.Encoder {
	return xmlEncoder{}
}
