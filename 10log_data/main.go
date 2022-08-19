package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

)

func get_time2() {
	now := time.Now()//获取当前时间
	fmt.Println(now)//输出时间格式：2022-05-17 12:19:47.011446 +0800 CST m=+0.006095901
	yearday := now.YearDay()//获取当前日期是一年中的第几天
	fmt.Printf("这是一年当中的第%d天\n",yearday)
	year := now.Year()//获取当前时间的年份
	month := now.Month()//获取当前时间的月份
	day := now.Day()//获取当前时间的日
	//获取时-分-秒
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func get_time() (msg string){
	now := time.Now()//获取当前时间
	timestamp := now.Unix()//时间戳
	timeObj := time.Unix(timestamp, 0)//将时间转化为时间格式
	fmt.Println(timeObj)//输出时间格式：2022-05-17 12:19:47 +0800 CST
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	msg = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	return msg
}
func log_err(str string)  {
	fmt.Println(str)
}
func log_info(str string)  {
	fileObj, err := os.OpenFile("./info.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log_err("log_info open info.txt failed.")
	}
	//关闭文件
	defer fileObj.Close()
	//填写日志信息
	msg := get_time() + str
	//使用bufio来写日志文件
	writer := bufio.NewWriter(fileObj)

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("caller failed!!!")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file)
	msg = msg + ":" + fileName + ":" + funcName
	msg = fmt.Sprintf("%s-%d", msg, line)
	writer.WriteString(msg)
	writer.WriteString("\n")
	writer.Flush()
}

func log_warning(str string) {
	fileObj, err := os.OpenFile("./warning.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log_err("log_warning open info.txt failed.")
	}
	//关闭文件
	defer fileObj.Close()
	//填写日志信息
	msg := get_time() + str
	//使用bufio来写日志文件
	writer := bufio.NewWriter(fileObj)
	writer.WriteString(msg)
	writer.WriteString("\n")
	writer.Flush()
}
func log_debug(str string)   {
	fileObj, err := os.OpenFile("./debug.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log_err("log_debug open info.txt failed.")
	}
	//关闭文件
	defer fileObj.Close()
	//填写日志信息
	msg := get_time() + str
	//使用bufio来写日志文件
	writer := bufio.NewWriter(fileObj)
	writer.WriteString(msg)//将数据写入到缓冲区
	writer.WriteString("\n")
	writer.Flush()//将缓存区中的内容写入到文件中
}
func main() {
	get_time2()
	log_info("未知登录")
	log_warning("出现警告")
	defer log_info("退出登录")
	log_debug("这是一条debug信息")
}
