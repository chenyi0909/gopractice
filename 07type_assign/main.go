package main

import "fmt"

//类型断言
func assign1(a interface{}){
	fmt.Printf("%T\n", a)
	str, ok := a.(string) //表示：猜测变量a的类型，str接收a的值，ok表示a的类型是否为string
	if !ok {
		fmt.Println("猜错了~")
	} else {
		fmt.Println("这是一个字符串：", str)
	}
}

func assign2(a interface{}){
	fmt.Printf("%T\n", a)
	switch  a.(type) {//类型断言：返回变量a的类型
	case string:
		fmt.Printf("%s 是一个 string 类型\n", a)
	case int:
		fmt.Printf("%d 是一个 int 类型\n", a)
	case bool:
		fmt.Printf("%d 是一个 bool 类型\n", a)
	default:
		fmt.Println("unknown type!")
	}
}

func main() {
	assign1("chenyi")
	assign2(true)
}
