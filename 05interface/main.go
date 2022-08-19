package main

import "fmt"

//同一个结构体可以实现多个接口
//接口嵌套
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat()
}

type cat struct {
	name string
}

func (c cat)move()  {
	fmt.Printf("(%s)猫走猫步~\n", c.name)
}
func (c cat)eat()  {
	fmt.Println("猫想吃鱼~")
}
func main() {
	//var nums1 []interface{}
	//nums2 := []int{1, 2, 3}
	//nums3 := append(nums1, nums2)
	//fmt.Println(len(nums1))
	var a []int
	b := append(a, 1)
	fmt.Println(b)
	//fmt.Println(len(nums2))
	//fmt.Println(len(nums3))
}
