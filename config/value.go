package config

import (
	"time"

	"git.code.oa.com/ad/go-project-base/config/reader"
)

type value struct{}

func newValue() reader.Value {
	return new(value)
}

// Bool ...
func (v *value) Bool(def bool) bool {
	return false
}

// Int ...
func (v *value) Int(def int) int {
	return 0
}

// String ...
func (v *value) String(def string) string {
	return ""
}

// Float64 ...
func (v *value) Float64(def float64) float64 {
	return 0.0
}

// Duration ...
func (v *value) Duration(def time.Duration) time.Duration {
	return time.Duration(0)
}

// StringSlice ...
func (v *value) StringSlice(def []string) []string {
	return nil
}

// StringMap ...
func (v *value) StringMap(def map[string]string) map[string]string {
	return map[string]string{}
}

// Scan ...
func (v *value) Scan(val interface{}) error {
	return nil
}

// Bytes ...
func (v *value) Bytes() []byte {
	return nil
}
