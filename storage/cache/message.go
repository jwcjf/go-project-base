package cache

import (
	"github.com/jwcjf/go-project-base/storage"
	"github.com/robinjoseph08/redisqueue/v2"
)

// Message ...
type Message struct {
	redisqueue.Message
}

// GetID ...
func (m *Message) GetID() string {
	return m.ID
}

// GetStream ...
func (m *Message) GetStream() string {
	return m.Stream
}

// GetValues ...
func (m *Message) GetValues() map[string]interface{} {
	return m.Values
}

// SetID ...
func (m *Message) SetID(id string) {
	m.ID = id
}

// SetStream ...
func (m *Message) SetStream(stream string) {
	m.Stream = stream
}

// SetValues ...
func (m *Message) SetValues(values map[string]interface{}) {
	m.Values = values
}

// GetPrefix ...
func (m *Message) GetPrefix() (prefix string) {
	if m.Values == nil {
		return
	}
	v, _ := m.Values[storage.PrefixKey]
	prefix, _ = v.(string)
	return
}

// SetPrefix ...
func (m *Message) SetPrefix(prefix string) {
	if m.Values == nil {
		m.Values = make(map[string]interface{})
	}
	m.Values[storage.PrefixKey] = prefix
}
