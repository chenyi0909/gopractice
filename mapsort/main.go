package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var scorMap = make(map[string]int, 20)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scorMap[key] = value
	}
	fmt.Println(scorMap)

	//
	var keys = make([]string, 0, 15)
	for key := range scorMap{
		keys = append(keys, key)
	}
	//
	sort.Strings(keys)
	//
	for _, key := range keys {
		fmt.Println(key, scorMap[key])
	}

}
/*
func main() {
	rand.Seed(time.Now().UnixNano())
	var scorMap = make(map[string]int, 20)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scorMap[key] = value
	}
	fmt.Println(scorMap)

	//
	var vs = make([]int, 0, 15)
	for _, v := range scorMap{
		vs = append(vs, v)
	}
	//
	sort.Ints(vs)
	//
	for _, v := range vs {
		fmt.Println(scorMap[key], v)
	}

}
 */