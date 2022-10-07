package main

import (
	"flag"
	"fmt"
)

func main() {
/*	fmt.Printf("%#v\n", os.Args)
	fmt.Println(os.Args[0], os.Args[1])
	fmt.Printf("%T\n", os.Args[0])*/

	//创建一个标志位
	//name := flag.String("name", "chenyi", "姓名")
	age := flag.Int("age", 25, "年龄")
	//
	var name string
	flag.StringVar(&name, "name", "chenyi", "姓名")
	//使用flag、解析参数
	flag.Parse()
	fmt.Println(name,*age)

	flag.Args()	//以切片形式返回命令行参数后的其他参数：29os.Args.exe -name 陈忆 -age 18 1 2 3（不包含程序名）
	flag.NArg()	//返回命令行参数后的其他参数的个数
	flag.NFlag()	//返回使用的命令行参数个数
}
