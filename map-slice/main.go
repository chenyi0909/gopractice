package main

import "fmt"

func main() {
	//元素类型为map的切片
	var s1 = make([]map[int]string, 1, 10)//make(切片类型，切片长度，切片容量)
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "chwe"

	fmt.Println(s1)

	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["beijing"] = []int{10,20,30}
	fmt.Println(m1)
}
