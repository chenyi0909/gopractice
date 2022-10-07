package main

import (
	"errors"
	"fmt"
	"os"
)

/*
使用数组实现队列的思路：
1.创建一个数组array，是作为队列的一个字段
2.front初始化为-1
3.rear：表示队列尾部，初始化为-1
4.完成队列的基本查找
	AddQueue ：加入数据到队列
	GetQueue ：从队列取出数据
	showQueue：显示队列
*/

const MaxSize = 4

// Queue 使用一个结构体管理队列
type Queue struct {
	array [MaxSize]int //数组==》模拟队列
	front int//表示指向队列首
	rear int//表示指向队列尾部
}

//AddQueue 添加数据到队列中
func (q *Queue) AddQueue(val int) (error) {
	//先判断队列已满
	if q.rear == MaxSize-1{//这个条件重要表示：rear指向队列尾部，包含队列尾部数据
		return errors.New("Queue is full.")
	}
	q.rear++
	q.array[q.rear] = val
	return nil
}

//ShowQueue 显示队列数据
func (q *Queue) ShowQueue() error {
	if q.front == q.rear{
		return errors.New("Queue is empty.")
	}
	for i, val := range q.array {
		fmt.Printf("array[%d]=%d\t",i,val)
		if i==q.rear{
			break
		}
	}
	fmt.Println("遍历结束")
	return nil
}

//GetQueue 从队列中取数据
func (q *Queue) GetQueue() (int,error) {
	if q.front == q.rear{
		return -1,errors.New("Queue is empty.")
	}
	q.front++
	val := q.array[q.front]
	return val,nil

}
func main() {
	//创建一个Queue结构体对象
	queue := &Queue{
		front: -1,
		rear: -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("Add failed,err:",err)
			}else{
				fmt.Println("add success")
			}
		case "show":
			err := queue.ShowQueue()
			if err != nil {
				fmt.Println("show queue failed,err:",err)
			}else {
				fmt.Println("show queue success")
			}
		case "get":
			ret, err := queue.GetQueue()
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
