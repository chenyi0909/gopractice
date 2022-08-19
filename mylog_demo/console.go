package mylog_demo

import (
	"fmt"
	"time"
)

//往终端输出日志
//Loggre 日志结构体
type Logger struct {
	Level LogLevel
}
//NewLog 构造函数
func NewLog(levelStr string) Logger{
	level,err := parseLogLevel(levelStr)
	if err != nil {
		//panic(err)
		fmt.Printf("parse log level failed, err:%v\n",err)
		//return
	}
	return Logger{
		Level: level,
	}
}

func (l Logger)enable(level LogLevel) bool {
	return level>=l.Level
}

func (l Logger)Debug(format string,a ...interface{}) {
		l.log(Debug,format, a...)
}

func (l Logger)Trace(format string,a ...interface{}) {
		l.log(Trace,format, a...)
}

func (l Logger)Info(format string,a ...interface{}) {
		l.log(Info,format, a...)
}

func (l Logger)Warning(format string,a ...interface{}) {
		l.log(Warning,format, a...)
}

func (l Logger)Error(format string,a ...interface{}) {
		l.log(Error,format, a...)
}

func (l Logger)Fatal(format string,a ...interface{}) {
		l.log(Fatal,format, a...)
}

func (l Logger)log(level LogLevel, format string, a ...interface{}) {
	if l.enable(level){
		msg := fmt.Sprintf(format, a...)
		funcName, fileName,line := getInfo(3)
		now := time.Now()
		fmt.Printf("[%s] [%s] [%s:%d:%s] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(level), fileName,line,funcName,msg)
	}
}


