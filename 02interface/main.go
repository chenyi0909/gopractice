package main

import "fmt"
//定义一个car类型的接口
type car interface {
	run()
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

func (f falali)run(){
	fmt.Printf("%s速度70迈~\n", f.brand)
}

func (b baoshijie)run()  {
	fmt.Printf("%s速度70迈~\n", b.brand)
}
//接收一个car类型的变量
func drive(car2 car) {
	car2.run()
}
func main() {
	var f = falali{
		"法拉利",
	}
	var b = baoshijie{
		"保时捷",
	}
	drive(f)
	drive(b)
}
