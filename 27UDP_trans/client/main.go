package main

import (
	"fmt"
	"net"
)

func main() {
	udpconn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("DialUDP failed, err:", err)
		return
	}
	defer udpconn.Close()
	sendData := []byte("ruanziyun")
	_, err = udpconn.Write(sendData)
	if err != nil {
		fmt.Println("sendUDP failed, err:", err)
		return
	}
	data := make([]byte, 1024)
	n, addr, err := udpconn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("readUDP failed:",err)
		return
	}
	fmt.Printf("From server[%v]:data:[%v],count:%v \n",addr,string(data[:n]), n)

}
