package main

import (
	"strings"
	"time"

	"github.com/baixiang1994/balog"
	"github.com/baixiang1994/banet"
)

func test_timeoutevent(connecter *banet.Connecter) {
	balog.Log("TimeoutEvent Test")
	connecter.SetReadTimeout(0)
	connecter.SetCallback(Handler)
}

func push_handler(connecter *banet.Connecter, length int, param interface{}) {
	balog.Info("push handler:", string(connecter.Buf))
}

func Handler(connecter *banet.Connecter, length int, param interface{}) {
	balog.Info("rcv:", string(connecter.Buf))
	cmd := strings.Split(string(connecter.Buf[:length]), " ")
	switch cmd[0] {
	case "PING":
		connecter.Write([]byte("PONG"))
	case "PUSH":
		if len(cmd) != 3 {
			connecter.Write([]byte("Error"))
		} else {
			balog.Log("its PUSH", cmd[1], cmd[2])
			connecter.SetReadTimeout(2 * time.Second)
			connecter.SetTimeoutEvent(test_timeoutevent)
			connecter.SetCallback(push_handler)
		}
	case "PULL":
		balog.Log("its PULL")
	default:
		balog.Info("Invalid command:", cmd[0])
		connecter.Write([]byte("Error Invalid"))
	}
}

func main() {
	balog.Log("Hello Blabo")

	server, err := banet.CreateServer(8080, "127.0.0.1")
	balog.RelAssert(err)
	server.SetCallback(Handler)
	server.SetBufSize(4096)
	go server.Start()
	select {}
}
