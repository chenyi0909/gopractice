package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子操作
var (
	x int64
	wg sync.WaitGroup
	//lock sync.Mutex
)

func add(){
	//lock.Lock()
	//x++
	//lock.Unlock()
	atomic.AddInt64(&x,1)//使用原子操作实现上方注释代码的功能
	wg.Done()
}
func main() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&x))
	//比较并交换
	ok := atomic.CompareAndSwapInt64(&x,100000,50)
	if ok == false {
		fmt.Println(atomic.LoadInt64(&x))
	} else{
		fmt.Println(atomic.LoadInt64(&x))
	}
}
