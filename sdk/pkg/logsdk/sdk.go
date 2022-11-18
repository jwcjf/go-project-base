package logsdk

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"git.code.oa.com/polaris/polaris-go/api"
)

// ProtoFlag proto flag
type ProtoFlag int32

// TCP ...
const (
	// TCP proto
	TCP = iota
	// UDP proto
	UDP = 1
)

const (
	nsProduction = "Production"
	nsTest       = "Test"
	svcTCP       = "zhiyan_log_collect_proxy_common_tcp"
	svcUDP       = "zhiyan_log_collect_proxy_common_udp"
)

// LogClient log client
type LogClient struct {
	conn       net.Conn
	proto      string
	protoFlag  ProtoFlag
	serverAddr string
	Topic      string
	Host       string
	lock       sync.Mutex

	usePolaris bool
	consumer   api.ConsumerAPI
	namespace  string
	service    string
	logChan    chan string
}

// InitConnect init connection
func (l *LogClient) InitConnect() error {
	var err error
	l.Close()

	if l.usePolaris {
		// use polaris to get ip and port
		l.serverAddr, err = l.getInstanceFromPolaris()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	// init connection
	conn, err := net.Dial(l.proto, l.serverAddr)
	if err != nil {
		fmt.Printf("addr: %s dail error. %s\n", l.serverAddr, err)
		return err
	}

	// if host is not set, use local addr
	if l.Host == "" {
		strs := strings.Split(conn.LocalAddr().String(), ":")
		if len(strs) > 0 {
			l.Host = strs[0]
		}
	}

	// just read, there is no reply from zhiyan now. 2020/03/17
	go func() {
		data := make([]byte, 1024)
		for {
			_, err := conn.Read(data)
			if err != nil {
				break
			}
		}
	}()

	l.conn = conn
	return nil
}

// Close close connection
func (l *LogClient) Close() {
	if l.conn != nil {
		l.conn.Close()
	}
}

// getInstanceFromPolaris get ip addr
func (l *LogClient) getInstanceFromPolaris() (string, error) {
	var err error
	if l.consumer == nil {
		l.consumer, err = api.NewConsumerAPI()
		if nil != err {
			fmt.Printf("fail to create ConsumerAPI by default configuration, err is %v\n", err)
			return "", err
		}
	}

	getInstancesReq := &api.GetOneInstanceRequest{}
	getInstancesReq.Namespace = l.namespace
	getInstancesReq.Service = l.service

	getInstResp, err := l.consumer.GetOneInstance(getInstancesReq)
	if err != nil {
		return "", err
	}

	if len(getInstResp.Instances) < 1 {
		return "", errors.New("no log server instance exist")
	}

	targetInstance := getInstResp.Instances[0]
	addr := fmt.Sprintf("%s:%d", targetInstance.GetHost(), targetInstance.GetPort())

	return addr, nil
}

// SetNameSpace set namespace
func (l *LogClient) SetNameSpace(namespace string) {
	l.namespace = namespace
}

// SendMessage send log message
func (l *LogClient) SendMessage(msg string) error {
	msgfmt := NewLogMessage(time.Now(), msg)
	pack := NewLogPack(l.Host, l.Topic, []*LogMessage{msgfmt})

	buf, err := pack.marshal()
	if err != nil {
		fmt.Println("log marshal error.", err)
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	l.InitConnect()
	if l.conn == nil {
		fmt.Println("logclient connection create error")
		return nil
	}

	totalLen := uint32(len(buf))
	if l.protoFlag == UDP {
		nLen, err := l.conn.Write(buf)
		if err != nil || nLen != int(totalLen) {
			fmt.Println("Write UDP error:", err)
			l.InitConnect()
			return errors.New("udp conncetion failed")
		}
	} else if l.protoFlag == TCP {
		var wroteLen uint32 = 0
		for {
			nLen, err := l.conn.Write(buf[wroteLen:totalLen])
			if err != nil {
				fmt.Println("Write TCP error:", err)
				l.InitConnect()
				return errors.New("tcp conncetion failed")
			}
			wroteLen += uint32(nLen)
			if wroteLen >= totalLen {
				break
			}
		}
	} else {
		fmt.Println("unknown protocal. ", l.proto)
	}

	return nil
}

// NewLogClient get a new logclient object,  if host is empty, we use localaddr
func NewLogClient(topic, proto, host string) *LogClient {
	proto = strings.ToLower(proto)

	logCli := &LogClient{Topic: topic, proto: proto, Host: host}

	if logCli.proto == "tcp" {
		logCli.protoFlag = TCP
		logCli.service = svcTCP
	} else {
		logCli.protoFlag = UDP
		logCli.service = svcUDP
	}
	logCli.SetNameSpace(nsProduction)
	logCli.usePolaris = true

	return logCli
}

// NewLogWithServerAddr Specific server addr ip and port manually
func NewLogWithServerAddr(serverAddr, topic, proto, host string) *LogClient {
	logCli := NewLogClient(topic, proto, host)
	logCli.usePolaris = false
	logCli.serverAddr = serverAddr

	return logCli
}
