package main

import "fmt"

//接口实现
type animal interface {
	move()
	eat()
}
type cat struct {
	name string
	feet int8 
}
type chicken struct {
	feet int8
}

func (c cat)move()  {
	fmt.Printf("(%s)猫走猫步~\n", c.name)
}
func (c cat)eat(){
	fmt.Println("猫想吃鱼~")
}
func (ch chicken)move()  {
	fmt.Println("鸡在走路~")
}
func (ch chicken)eat()  {
	fmt.Println("鸡吃饲料~")
}

func handle(a animal)  {
	a.move()
	a.eat()
}

func main() {
	//var a1 animal
	var c = cat{
		"月饼",
		4,
	}
	handle(c)

}
