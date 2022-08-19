package main

import (
	mylog_demoaa "goprac/18chan_log_demo"
	"time"
)

func main(){
/*		var log mylog_demoaa.Methoder//使用接口变量
		log = mylog_demoaa.NewLog("Debug")//往终端输出日志
		id := 10068
		name := "chenyi"
		log.Debug("id:%d-name:%s 这是一条Debug日志", id, name)
		log.Trace("这是一条Trace日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")*/
	var log mylog_demoaa.Methoder//使用接口变量
	log = mylog_demoaa.NewFileLogger("Debug","./","chenyi.log", 10*1024)//往文件写入日志
	id := 10068
	name := "chenyi"
	for  {
		log.Debug("id:%d-name:%s 这是一条Debug日志", id, name)
		log.Trace("这是一条Trace日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}