package main

import (
	"fmt"
	"net"
)

func readConn(conn net.Conn)  {
	//3.与客户端通信
	var tmp [128]byte
	for  {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read err:",err)
			return
		}
		fmt.Println(string(tmp[:n]))
		conn.Write([]byte("收到！"))
	}
}
func main() {
	//1.启动本地端口服务
	listener, err := net.Listen("tcp","127.0.0.1:5060")
	if err != nil {
		fmt.Println("socket err:",err)
		return
	}
	//2.等待建立连接
	for  {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:",err)
			return
		}
		go readConn(conn)
	}
}
