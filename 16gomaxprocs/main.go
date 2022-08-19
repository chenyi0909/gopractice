package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
func a(){
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("a:",i)
	}
}

func b()  {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("b:",i*i)
	}
}
func main() {
	runtime.GOMAXPROCS(6)//默认程序跑满所有cpu
	fmt.Println("numcpu:",runtime.NumCPU())//获取服务器cpu个数
	wg.Add(1)
	go a()
	wg.Add(1)
	go b()
	wg.Wait()
}
