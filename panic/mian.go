package main

import (
	"fmt"
	"log"
)

func funcA()  {

	fmt.Println("a")
}
func funcB()  {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放资源...")
	}()
	log.Panicln("find serial err")
	fmt.Println("b")
}
func funcC()  {
	fmt.Println("c")
}
func main() {
	funcA()
	funcB()
	funcC()
}
