package main

import "fmt"

func main() {
	array := [...]int{1,2,2,5,4,6,3,7,1,0}
	fmt.Println(array)
	slice := array[:5]
	fmt.Println(slice)
	array[3] = 10
	array[2] = 55
	fmt.Println(array)
	fmt.Println(slice)
	//切片是引用类型
}
