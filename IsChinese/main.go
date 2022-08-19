package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "Hello世界"
	count := 0
	for _, c := range s1 {
		if unicode.Is(unicode.Han,c){
			count++
		}
	}
	fmt.Println(count)

	s2 := "how do you do"
	s3 := strings.Split(s2," ")
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		if _,ok := m1[w]; !ok{
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//回文判断
	ss := "成都是都城"
	fmt.Println("len(ss): ", len(ss))
	//把字符串中的字符拿出来放到一个[]rune中
	r := make([]rune,0,len(ss))
	for _, c := range ss {
		r = append(r,c)
	}
	fmt.Println(r)
	fmt.Println("len(r): ", len(r))
	//
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i]{
			fmt.Println("字符串不是回文")
			return
		}
	}
	fmt.Println("字符串是回文")
}
