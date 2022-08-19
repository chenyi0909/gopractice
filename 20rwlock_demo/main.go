package main

import (
	"fmt"
	"sync"
	"time"
)

//rwlock
var(
	x = 0
	rwlock sync.RWMutex //读写锁
	lock sync.Mutex
	wg sync.WaitGroup
	once sync.Once
)
func read(){
	rwlock.RLock()//加读锁
	fmt.Println(x)
	//time.Sleep(time.Millisecond*500)
	rwlock.RUnlock()//释放读锁
	wg.Done()
}

func write(){
	rwlock.Lock()//加写锁：其他goroutine无法在进行读写操作
	x = x+1
	//time.Sleep(time.Millisecond*500)
	rwlock.Unlock()//释放写锁
	wg.Done()
}
func main() {
	start := time.Now()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go write()
	}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go read()
	}
	once.Do()
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
