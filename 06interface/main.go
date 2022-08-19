package main

import "fmt"

//定义一个空接口
//type any interface {}

func main() {
	//var a any
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "chenyi"
	m1["age"] = 18
	m1["merried"] = false
	m1["hobby"] = [...]string{"qq", "qww", "sds"}
	fmt.Println(m1)
}
