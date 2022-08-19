package main

import "fmt"

//闭包
func f1(f func())  {
	fmt.Println("this is f1")
	f()
}
func f2(x,y int)  {
	fmt.Println("this is f2")
	fmt.Println(x+y)
}
//要求：
//将f2当作参数传入f1  f1(f2)
func f3(f func(int,int),x, y int) func() {
	tmp := func() {
		//fmt.Println(x+y)
		f(x,y)
	}
	return tmp
}
func main() {
	ret:=f3(f2, 11, 22)
	f1(ret)
}
