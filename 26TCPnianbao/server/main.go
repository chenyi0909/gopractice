package main

import (
	"bufio"
	"fmt"
	proto "goprac/protocol"
	"io"
	"net"
)

func readConn(conn net.Conn)  {
	defer conn.Close()
	//3.与客户端通信
	reader := bufio.NewReader(conn)
	for  {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			fmt.Println("read EOF:",err)
			return
		}
		if err != nil {
			fmt.Println("decode err:",err)
			return
		}
		fmt.Println(msg)
		conn.Write([]byte("收到！"))
	}
}
func main() {
	//1.启动本地端口服务
	listener, err := net.Listen("tcp","127.0.0.1:30000")
	defer listener.Close()
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
