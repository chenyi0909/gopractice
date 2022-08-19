package main

import "fmt"

//阶乘递归
func fac(n int) int {
	if n<= 1 {
		return 1
	}
	return n*fac(n-1)
}

//上台阶面试题
func raiseFloot(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return raiseFloot(n-1)+raiseFloot(n-2)
}
func main() {
	//var a int
	//fmt.Print("输入要计算的阶乘：")
	//fmt.Scanf("%d", &a)
	//ret := fac(a)
	//fmt.Println(ret)
	var n uint64
	fmt.Print("输入要上的台阶数：")
	fmt.Scanf("%d", &n)
	ret := raiseFloot(n)
	fmt.Printf("一共有%d种上台阶的方法。", ret)
}
