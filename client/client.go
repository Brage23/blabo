package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/baixiang1994/banet"

	"github.com/baixiang1994/balog"
)

func main() {
	balog.Log("Hello Banet Client")
	if len(os.Args) != 3 {
		balog.Log("input param invalid")
		return
	}

	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		balog.Log("param port invalid")
		return
	}
	balog.Info("input ADDR ->", os.Args[1], ":", port)

	client, err := banet.CreateClient(port, os.Args[1])
	if err != nil {
		balog.Log("connect to host failed:", err)
		return
	}

	var command string
	for {
		rcvbuf := make([]byte, 64)
		fmt.Scanln(&command)
		balog.Info("scan:", command)
		cmd := strings.Split(command, " ")
		switch cmd[0] {
		case "ping":
			client.Write([]byte("PING"))
			_, err = client.ReadwithTimeout(rcvbuf, 2*time.Second)
			if err != nil {
				balog.Error("io timeout")
			} else {
				fmt.Println(string(rcvbuf))
			}
		default:
			balog.Info("Invalid command:", cmd[0])
		}
	}
	/*
		client.Write([]byte("PUSH /home/bai/test.log 1024"))
		time.Sleep(1 * time.Second)
		client.Write([]byte("Hi"))
		time.Sleep(3 * time.Second)
		client.Write([]byte("PING"))
		rcvbuf := make([]byte, 32)
		len, err := client.Read(rcvbuf)
		balog.Info("ClientRcv:", string(rcvbuf[:len]))
		_, err = client.ReadwithTimeout(rcvbuf, 2*time.Second)
		balog.Log(err)
		client.Read(rcvbuf)
		balog.Log("Read done")
	*/
	client.Close()
}
