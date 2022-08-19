package main

import "fmt"

//函数
//函数定义
//有返回值
func sum(x int, y int) int  {
	return x+y

}
//没有返回值
func f1(s int, a int)  {
	fmt.Println(s+a)
}

//没有参数和返回值
func f2()  {
	fmt.Println("f2 没有返回值")
}

//没有参数但有返回值
func f3() int {
	ret := 3
	fmt.Println("没有参数但有返回值 ret= ", ret)
	return ret
}

//返回值也可以是命名
func f4(x int, y int) (ret int) {
	ret = x+y
	fmt.Println("返回值是命名： ret = ", ret)
	return  ret;
}

//多个返回值
func f5() (int, string) {

	fmt.Println("多个返回值")
	return 1, "cjemyuoi"

}

//参数简写
func f6(x, y int, s, w string)  {
	fmt.Println("参数简写")
}

//可变长参数
//可变长参数必须放在参数列表最后
func f7(x string, y ...int)  {
	fmt.Println(x)
	fmt.Printf("%T", y)// y是切片类型
	fmt.Println("可变长参数")
}
func main() {
	r := sum(1, 2)
	fmt.Println(r)
}
