package mylog_demo

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志
//定义一个日志对象结构体
type FileLogger struct {
	level LogLevel
	filePath string
	fileName string
	maxFileSize int64
	fileObj *os.File
	errFileObj *os.File
}

//日志对象的构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		level: logLevel,
		filePath: fp,
		fileName: fn,
		maxFileSize: maxSize,
	}
	f1.initFile()
	return f1
}

func (f *FileLogger)initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}
//判断是否需要记录日志
func (f *FileLogger)enable(level LogLevel) bool {
	return level>=f.level
}
//检查日志大小，判断是否需要切割文件
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file state failed, err:%v\n",err)
		return false
	}
	//如果当前文件的大小大于等于要写入文件的大小的最大值，返回true，否则返回false
	return fileInfo.Size() >= f.maxFileSize

}
//q切割日志文件
func (f *FileLogger) sqliteFileBySize(fileObj *os.File) (*os.File,error) {
	fileName := f.fileObj.Name()
	fmt.Println(fileName)
	//1、关闭当前日志文件
	f.fileObj.Close()
	//2、备份一下 rename要被切割的日志文件
	nowStr := time.Now().Format("20060102150405000")
	logName := fmt.Sprintf("%s\\%s", f.filePath,fileName)
	newlogName := fmt.Sprintf("%s\\%s.bak%s", f.filePath, fileName, nowStr)
	os.Rename(logName, newlogName)//将原来的日志文件重命名
	//3、打开一个新的日志文件
	newFileObj, err:= os.OpenFile(logName,os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n",err)
		return nil,errors.New("open new log file failed")
	}
	//4、将打开的新的日志文件对象赋值给f.fileObj --> 这一步改在调用函数返回之后执行
	return newFileObj,nil
}
//记录日志的方法
func (f *FileLogger)log(level LogLevel, format string, a ...interface{}) {
	if f.enable(level) {
		msg := fmt.Sprintf(format, a...)
		funcName, fileName,line := getInfo(3)
		now := time.Now()
		if f.checkSize(f.fileObj) {
			//此时需要切割当前日志文件
			fileObj, err := f.sqliteFileBySize(f.fileObj)
			if err != nil {
				fmt.Println("sqliteFileBySize() return err")
				return
			}
			f.fileObj = fileObj
		}
		fmt.Fprintf(f.fileObj,"[%s] [%s] [%s-%d-%s] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(level), fileName,line,funcName,msg)
		if level >= Error{
			//如果记录日志的级别大于等于Error的话，还需要将日志记录于err日志文件中
			if f.checkSize(f.errFileObj) {
				//此时需要切割当前日志文件
				errFileObj, err := f.sqliteFileBySize(f.errFileObj)
				if err != nil {
					fmt.Println("sqliteFileBySize(errFileObj) return err")
				}
				f.errFileObj = errFileObj
			}
			fmt.Fprintf(f.errFileObj,"[%s] [%s] [%s-%d-%s] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(level), fileName,line,funcName,msg)
		}
	}
}

func (f *FileLogger)Debug(format string,a ...interface{}) {
		f.log(Debug,format, a...)
}

func (f *FileLogger)Trace(format string,a ...interface{}) {
		f.log(Trace,format, a...)
}

func (f *FileLogger)Info(format string,a ...interface{}) {
		f.log(Info,format, a...)
}
func (f *FileLogger)Warning(format string,a ...interface{}) {
		f.log(Warning,format, a...)
}

func (f *FileLogger)Error(format string,a ...interface{}) {
		f.log(Error,format, a...)
}

func (f *FileLogger)Fatal(format string,a ...interface{}) {
		f.log(Fatal,format, a...)
}

