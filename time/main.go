package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Date())//显示年 月 日
	fmt.Println(now.Zone())
	//时间戳
	t := now.Unix()
	//j将时间戳转换为时间格式
	ret := time.Unix(t, 0)
	fmt.Println(ret)
	fmt.Println(ret.Date())//显示年 月 日
	//now + 1h
	fmt.Println(now.Add(24*time.Hour))
	//定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t)//每一秒钟执行一次
	//}
	//格式化时间 把语言中的时间对象 转换成字符串类型的时间
	//fmt.Println(now.Format("2006-01-02"))
	//fmt.Println(now.Format("2006-01-02 15:04:05"))//24进制
	//fmt.Println(now.Format("2006-01-02 03:04:05 PM"))//12进制
	//sub 两个时间相减
	timeObj, err := time.Parse("2006-01-02", "2022-06-02")//将参数2按照参数1的格式解析，并返回时间对象
	if err != nil {
		fmt.Println("parse time err, err:", err)
		return
	}
	fmt.Println(timeObj)//timeObj为零时区的时间，我国时间比0时区时间早8个小时
	d := now.Sub(timeObj)//时间差：now-timeObj
	fmt.Println(d)
	//关于Sleep
	//n := 5
	//time.Sleep(time.Duration(n)*time.Second)//如果传参为变量，必须显式转换为Duration类型才能使用
	//time.Sleep(5*time.Second)//如果传参为数字，隐式转换为Duration类型
	e := timeObj.Sub(now)//返回时间间隔：timeObj-now
	f := timeObj.Sub(now.Add(8*time.Hour))//相当于timeObj-now-8hour：因为当前时区比0时区早8小时
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("---------------------------------------")
	//按照东八区的时区和格式解析字符串格式
	//1、根据字符串加载时区
	//var loc *time.Location
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load faild")
		return
	}
	//2、按照指定时间解析
	timeObj, err = time.ParseInLocation("2006-01-02", "2022-06-02", loc)
	if err != nil {
		fmt.Println("parse failed")
		return
	}
	fmt.Println(timeObj)
	g := timeObj.Sub(now)//返回时间间隔：timeObj-now
	fmt.Println(g)

}
