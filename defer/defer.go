package main

import "fmt"

func deferDemo()  {
	fmt.Println("chen")
	defer fmt.Println("shuai")//一个函数中可以有多个defer，执行顺序按照先进后出的顺序
	defer fmt.Println("zhen")
	fmt.Println("yi")
	//如果函数有返回值，则defer语句执行在return语句之间
	//
}
func main() {
	deferDemo()
}
