package logsdk

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mashalbytes = []byte{
	9, 1, 2, 3,
	0, 0, 0, 46,
	2,
	9, 49, 50, 55, 46, 48, 46, 48, 46, 49,
	7, 116, 101, 115, 116, 48, 48, 49,
	0,
	0, 0, 1, 111, 94, 102, 232, 0, // 20200101 UTC
	0, 0, 0, 10,
	104, 101, 108, 108, 111, 119, 111, 114, 108, 100,
}

var host string = "127.0.0.1"
var topic string = "test001"
var msg string = "helloworld"

func TestPkgMarShal(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	ts := time.Date(2020, 1, 1, 0, 0, 0, 0, loc)

	msgfmt := NewLogMessage(ts, msg)

	pack := NewLogPack(host, topic, []*LogMessage{msgfmt})

	buf, _ := pack.marshal()

	assert.Equal(t, mashalbytes, buf)

}

func BenchmarkPkgMarShal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgfmt := NewLogMessage(time.Now(), msg)
		pack := NewLogPack(host, topic, []*LogMessage{msgfmt})

		_ = pack
	}
}
