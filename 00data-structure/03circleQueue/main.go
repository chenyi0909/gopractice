package main

import (
	"errors"
	"fmt"
	"os"
)

const maxSize = 4
//CircleQueue 使用一个结构体管理数组
type CircleQueue struct {
	arrary [maxSize]int
	head int//指向队列队首
	tail int//指向队列队尾,队列的最后一个元素的后一个位置
}
//isFull 判断队列是否已满
func (cq *CircleQueue) isFull() bool {
	if (cq.tail+1)%maxSize == cq.head{
		return true
	}
	return false
}
//isEmpty 判断队列是否为空
func (cq *CircleQueue) isEmpty() bool {
	if cq.head == cq.tail{
		return true
	}
	return false
}
//QueueSize 取出队列已有多少元素
func (cq *CircleQueue) QueueSize() int {
	return (cq.tail+maxSize-cq.head)%maxSize //计算队列大小
}
//AddQueue 添加队列元素
func (cq *CircleQueue) AddQueue(val int) error {
	if cq.isFull(){
		return errors.New("the CircleQueue is full, add failed")
	}
	//cq.tail指向队列的最后一个元素的后一个位置
	cq.arrary[cq.tail] = val
	/*if cq.tail == maxSize-1{
		cq.tail = 0
	}else {
		cq.tail++
	}*/
	cq.tail = (cq.tail + 1)%maxSize
	return nil
}
//GetQueue 从队列中取出数据
func (cq *CircleQueue) GetQueue() (int,error)  {
	if cq.isEmpty(){
		return 0,errors.New("the queue is Empty, no data in the queue")
	}
	//取出队列元素
	val := cq.arrary[cq.head]
	/*if cq.head == maxSize-1{
		cq.head = 0
	}else {
		cq.head++
	}*/
	cq.head = (cq.head + 1)%maxSize
	return val,nil
}
func (cq *CircleQueue) ShowQueue() error {
	if cq.isEmpty(){
		return errors.New("the queue is Empty, no data in the queue")
	}
	//取出队列元素个数
	size := cq.QueueSize()
	//定义一个临时变量
	tempHead := cq.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t",tempHead,cq.arrary[tempHead])
		tempHead = (tempHead + 1)%maxSize
	}
	fmt.Println()
	return nil
}
func main() {
	//生成一个环状队列对象
	circlequeue := &CircleQueue{
		head: 0,
		tail: 0,
	}
	var key string
	var val int
	for  {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Printf("请输入数据：")
			fmt.Scanln(&val)
			err := circlequeue.AddQueue(val)
			if err != nil {
				fmt.Println("Add failed,err:",err)
			}else{
				fmt.Println("add success")
			}
		case "show":
			err := circlequeue.ShowQueue()
			if err != nil {
				fmt.Println("show queue failed,err:",err)
			}else {
				fmt.Println("show queue success")
			}
		case "get":
			ret, err := circlequeue.GetQueue()
			if err != nil {
				fmt.Println("get value failed,err:",err)
			}else {
				fmt.Printf("get value success, the value : [%d]\n", ret)
			}
		case "exit":
			fmt.Println("exit...")
			os.Exit(0)
		}
	}
}
