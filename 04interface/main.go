package main

import "fmt"

//使用值接收者和指针接收者的区别？
type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}
//值接收者
func (c cat)move()  {
	fmt.Printf("(%s)猫走猫步~\n", c.name)
}
//指针接收者
//func (c *cat)move()  {
//	fmt.Printf("(%s)猫走猫步~\n", c.name)
//}
func (c cat)eat(s string ){
	fmt.Printf("猫想吃%s~\n", s)
}
func main() {
	var a1 animal
	var a2 animal

	c1 := cat{"Tom",4}
	c2 := &cat{"假老练", 4}

	a1 = c1
	a2 = c2
	fmt.Println(a1)
	a1.move()
	a2.move()
}
