package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 30000,
	})
	//listen, err := net.Listen("udp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listenUDP failed, err:", err)
		return
	}
	defer listen.Close()
	for  {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("readFromUDP failed, err:", err)
			continue
		}
		fmt.Printf("addr:%v,data:%v,count:%v\n", addr,string(data[:n]),n)
		_, err = listen.WriteToUDP([]byte("服务端已收到。"), addr)
		if err != nil {
			fmt.Println("write udp failed, err:", err)
			continue
		}
	}
}
