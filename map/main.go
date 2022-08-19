package main

import "fmt"

func main() {
	m1 := make(map[string]int, 10)

	m1["chenyi"] = 22
	m1["xiaoli"] = 21
	m1["ruan"] = 20
	fmt.Println(m1)

	//通过key查找value
	value, ok := m1["chenyi"]
	if !ok{
		fmt.Println("no people")
	}else {
		fmt.Println(value)
	}

	//遍历
	for key := range m1 {
		fmt.Println(key)
	}
	for _, v := range m1 {
		fmt.Println(v)

	}

	for k, v := range m1 {
		if v == 21{
			delete(m1,k)
		}
	}

	//删除元素
	//delete(m1,"xiaoli")
	fmt.Println(m1)
	delete(m1,"chj")
	fmt.Println(m1)


}

