package main

import (
	"agent/conf"
	"agent/tcp"
)

func main() {
	go conf.HandleConf()
	go tcp.TcpConnect()
	select {}
}
