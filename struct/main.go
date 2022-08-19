package main

import "fmt"

//结构体

type person struct {
	name string
	age int
	gender string
	hobby []string
}
func main() {
	var  p person
	//匿名结构体
	var s struct{
		x int
		y int
	}
	p.name = "chenyi"
	p.age = 20
	p.gender = "男"
	p.hobby = []string{"篮球", "滑板"}
	fmt.Println(p)
	fmt.Printf("%T\n",p)
	fmt.Println(p.name)


	s.x = 1
	s.y = 2
}
