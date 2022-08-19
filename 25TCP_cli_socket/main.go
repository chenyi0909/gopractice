package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//1.与服务端建立连接
	conn,err := net.Dial("tcp","127.0.0.1:5060")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	//2.发送数据
	//var msg string//用于存储客户端发送的消息
	var ret [128]byte//用于接收服务端返回的消息
	reader := bufio.NewReader(os.Stdin)
	for  {
		fmt.Printf("请输入：")
		//fmt.Scanln(&msg)//遇到空格就输入结束
		msg,_:=reader.ReadString('\n')//该方法表示读到换行符结束
		fmt.Printf("%#v\n", msg)
		if msg == "exit\r\n" {//使用\r\n 的原因是，linux换行符是\r\n
			fmt.Println("退出....")
			break
		}
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("发送err:",err,"请重新建立连接...")
		}
		n,_ := conn.Read(ret[:])
		fmt.Println("来自服务端：",string(ret[:n]))
	}
	conn.Close()
}
