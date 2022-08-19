package main

import "fmt"

//闭包是一个函数，这个函数包含了一个外部变量
func adder(x int) func(int) int {
	//x为环境变量，该变量不会随着调用结束而释放，
	//会一直存在直到程序结束（也就是闭包会在程序结束才结束）
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	ret := adder(100)
	ret2 := ret(200)
	fmt.Println(ret2)
}
