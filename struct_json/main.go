package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json
//1.序列化：把Go语言中的结构体变量 --> json格式的字符串
//2.反序列化：json格式的字符串  --> Go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age int `json:"age"`
}
func main() {
	p1 := person{
		Name:"chenyi",
		Age:22,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil{
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))
	//反序列化
	str := `{"name":"xiaoli","age":20}`
	var p2 person//用于接收反序列化字符串str的值
	json.Unmarshal([]byte(str), &p2)//传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}
