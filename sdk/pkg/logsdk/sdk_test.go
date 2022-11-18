package logsdk

import (
	"testing"
)

func TestSDK(t *testing.T) {
	topic := "fb-52ec854gc4eff42"
	// topic = "fb-1ddf7b977e4eag7"

	myIP := ""
	logCli := NewLogClient(topic, "udp", myIP)
	logCli.SetNameSpace(nsProduction)
	logCli.InitConnect()

	logCli.SendMessage("test log 123456 udp")
}

func TestSDKWithAddr(t *testing.T) {
	topic := "fb-52ec854gc4eff42"
	serverAddr := ""
	// serverAddr := "127.0.0.1:11001"

	logCli := NewLogWithServerAddr(serverAddr, topic, "tcp", "")
	logCli.InitConnect()

	logCli.SendMessage("test log 123456 udp")
}
