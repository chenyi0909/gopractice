package main

import "fmt"

type cat struct {}

type dog struct {}

//定义了一个叫做speaker的接口
type speaker interface {
	speak()
}
//cat实现方法
func (c cat)speak(){
	fmt.Println("喵喵喵~")
}
//dog实现方法
func (d dog)speak(){
	fmt.Println("汪汪汪~")
}
func da(x speaker)  {
	//接收一个参数，传进来什么 就打什么
	x.speak()//被打的就要叫唤
}
func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)


}
