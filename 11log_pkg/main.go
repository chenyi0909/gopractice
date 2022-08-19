package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileObj, err := os.OpenFile("./debug.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Sprintf("open debug file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	log.SetOutput(fileObj)//定义日志往文件里面写
	log.Println("普通日志")

}
