package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//1
	//fileObj, err := os.Open("./main.go")
	//if err != nil{
	//	fmt.Printf("open file failed, err:%v\n", err)
	//	return
	//}
	//defer fileObj.Close()
	////
	//var temp = make([]byte, 128)
	//for  {
	//	n, err := fileObj.Read(temp[:])
	//	if err != nil {
	//		fmt.Printf("read from file failed, err:%v\n", err)
	//		return
	//	}
	//	fmt.Printf("当前读了%d个字节\n", n)
	//	fmt.Println(string(temp[:n]))
	//	if n <= 0{
	//		return
	//	}
	//}
	//readFromFileBufio()
	readFromFileByIoutil()
}

//2
func readFromFileBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil{
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	//
	reader := bufio.NewReader(fileObj)
	for  {
		line, err := reader.ReadString('\n')
		if err==io.EOF {
			fmt.Println("文件读取完毕！！！")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		fmt.Println(line)
	}
}
//3
func readFromFileByIoutil(){

	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read from file failed, err:%v\n", err)
	}
	fmt.Println(string(ret))
}