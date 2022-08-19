package main

import (
	"fmt"
	"strings"
)
func main() {
	s := "hello chenyi"
	fmt.Println(s)
	n := len(s)
	fmt.Println(n)

	for i := 0; i < len(s); i++{
		fmt.Printf("%c", s[i])
	}
	fmt.Println()

	for _, c := range s {
		if c != '\n' {
			fmt.Printf("%c", c)
		}
	}
	s1 := strings.Clone(s)

	fmt.Println("copy s->s1 : ", s1)

	fmt.Println(strings.Split(s1,"h") )

	fmt.Println(strings.Join(strings.Split(s1,"h"), "+"))
}
