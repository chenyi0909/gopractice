package mylog_demoaa

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//自定义一个日志库

type LogLevel uint16

type Methoder interface {
	Debug(format string,a ...interface{})
	Trace(format string,a ...interface{})
	Info(format string,a ...interface{})
	Warning(format string,a ...interface{})
	Error(format string,a ...interface{})
	Fatal(format string,a ...interface{})
}

const (
	Debug LogLevel= iota
	Trace
	Info
	Warning
	Error
	Fatal
	Unknown
)

func parseLogLevel(s string) (LogLevel,error)  {
	s = strings.ToLower(s)
	switch s{
	case "debug":
		return Debug,nil
	case "trace":
		return Trace,nil
	case "info":
		return Info,nil
	case "warning":
		return Warning,nil
	case "error":
		return Error,nil
	case "fatal":
		return Fatal,nil
	default:
		return Unknown, errors.New("无效的日志级别")

	}
}

func getInfo(n int) (funcName, fileName string, lineNo int) {
	pc, file, line, ok:=runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller failed")
	}
	lineNo = line
	funcName = runtime.FuncForPC(pc).Name()//
	fileName = path.Base(file)//
	return funcName, fileName, lineNo
}

func getLogString(lv LogLevel) string {
	switch lv {
	case Debug:
		return "DEBUG"
	case Trace:
		return "TRACE"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		return "DEBUG"
	}
}