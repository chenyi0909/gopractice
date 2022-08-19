package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//fileObj, err := os.OpenFile("./aa.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	//if err != nil{
	//	fmt.Printf("open file failed, err:%v\n", err)
	//	return
	//}
	//defer fileObj.Close()
	////Write
	//fileObj.Write([]byte("chenyi and ruanziyun\n"))
	////WriteString
	//fileObj.WriteString("北京欢迎你!!!")

	writeToFileBufio()
	writeToFileByIoutil()
}

func writeToFileBufio()  {
	fileObj, err := os.OpenFile("./aa.txt", os.O_TRUNC|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil{
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	//创建一个写对象
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("成都")//将数据写入到缓存区
	writer.Flush()//将缓存区中的内容写入到文件中
}
func writeToFileByIoutil() {
	str := "，一座来了就不想走的城市。"
	err := ioutil.WriteFile("./aa.txt",[]byte(str), 0644 )
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
}