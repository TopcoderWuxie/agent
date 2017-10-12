package tcp

import (
	"net"

	. "agent/conf"
)

func TcpConnect() {
	var tcpListener *net.TCPListener
	for {
		var err error
		localAddress, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:8088") //定义一个本机IP和端口
		tcpListener, err = net.ListenTCP("tcp", localAddress)         //在刚定义好的地址上进监听请求
		if err != nil {
			DebugLog.Println("listent tcp is failed, the reason is:", err)
			tcpListener.Close()
		} else {
			break
		}
	}

	for {
		conn, err := tcpListener.Accept() //接受连接
		if err != nil {
			DebugLog.Println("tcp connect is error")
			continue
		}

		go Handle()
		conn.Close()
	}
}
