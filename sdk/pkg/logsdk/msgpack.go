package logsdk

import (
	"encoding/binary"
	"time"
)

const (
	magicHead  uint32 = 0x09010203
	packMinLen uint32 = 4 + 1 + 1 + 1 + 1
	msgMinLen  uint32 = 8 + 4
)

// LogMessage zhiyan log message define
type LogMessage struct {
	Timestamp  uint64
	MessageLen uint32
	Message    []byte
}

// LogPack zhiyan log package define
type LogPack struct {
	PkgLen   uint32        // 包总长度 (标识符不计算在内)
	PkgVer   uint8         // 包版本号，0x02
	HostLen  uint8         // host 长度
	Host     []byte        // host 字符串
	TopicLen uint8         // topic 长度
	Topic    []byte        // topic 字符串
	Compress uint8         // 压缩选项，0x00 无，0x01 gzip
	Messages []*LogMessage // 消息体（若 compress!=0x00 则为压缩后的 buffer）
}

func (p *LogPack) marshal() ([]byte, error) {
	p.PkgLen = packMinLen
	p.PkgVer = 0x02

	for _, v := range p.Messages {
		p.PkgLen += msgMinLen
		p.PkgLen += uint32(len(v.Message))
	}

	p.PkgLen += uint32(p.HostLen + p.TopicLen)

	buf := make([]byte, p.PkgLen+4)
	pos := 0

	binary.BigEndian.PutUint32(buf[pos:], magicHead)
	pos += 4
	binary.BigEndian.PutUint32(buf[pos:], p.PkgLen)
	pos += 4
	buf[pos] = p.PkgVer
	pos++
	buf[pos] = p.HostLen
	pos++
	copy(buf[pos:], p.Host)
	pos += len(p.Host)

	buf[pos] = p.TopicLen
	pos++
	copy(buf[pos:], p.Topic)
	pos += len(p.Topic)

	buf[pos] = p.Compress
	pos++

	for _, v := range p.Messages {
		binary.BigEndian.PutUint64(buf[pos:], v.Timestamp)
		pos += 8
		binary.BigEndian.PutUint32(buf[pos:], v.MessageLen)
		pos += 4
		copy(buf[pos:], v.Message)
		pos += len(v.Message)
	}

	return buf, nil
}

// NewLogMessage get a new log message object
func NewLogMessage(t time.Time, msg string) *LogMessage {
	logmsg := &LogMessage{
		Timestamp:  uint64(t.UnixNano() / 1000000),
		MessageLen: uint32(len(msg)),
		Message:    []byte(msg),
	}

	return logmsg
}

// NewLogPack get a new logpack object
func NewLogPack(host, topic string, msgs []*LogMessage) *LogPack {
	pack := &LogPack{
		HostLen:  uint8(len(host)),
		Host:     []byte(host),
		TopicLen: uint8(len(topic)),
		Topic:    []byte(topic),
		Compress: 0x00,
		Messages: msgs,
	}

	return pack
}