package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ValNode struct {
	row int
	col int
	val int
}
func main() {
	//1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1//黑子
	chessMap[2][3] = 2//蓝子
	
	//2.输出看看原始数组
/*	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			fmt.Printf("%d\t", chessMap[i][j])
		}
		fmt.Println()
	}*/
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

	//3.转成稀疏数组。
	//思路
	//（1）遍历chessMap，如果我们发现有一个元素的值不为0，就创建一个node结构体
	//（2）将其放入到对应的切片即可

	var storeStr string//将存盘的数据输出未一个字符串并保存到文件中
	var sparseArr []ValNode

	//标准的一个稀疏数组应该还有一个记录元素的二维数组的规模（行、列和值）
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0{
				//创建一个ValNode值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//输出稀疏数组
	fileObj,err := os.OpenFile("./chessMap.data",os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("文件打开错误，存盘失败...err:",err)
		return
	}
	//关闭存盘文件
	//defer fileObj.Close()
	fmt.Println("当前的稀疏数组：")
	for i, valNode := range sparseArr {
		fmt.Printf("%d : %d, %d, %d\n", i, valNode.row,valNode.col,valNode.val)
		storeStr = fmt.Sprintf("%d\t%d\t%d\n", valNode.row,valNode.col,valNode.val)
		_, err = fileObj.WriteString(storeStr)
		if err != nil {
			fmt.Println("存盘失败...")
			return
		}
	}
	//关闭存盘文件
	fileObj.Close()

	//4.恢复原始数组，从存盘文件中读取数据并恢复存盘数组
	fileObj, err = os.OpenFile("./chessMap.data",os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("读取存盘数据失败...err:", err)
		return
	}
	defer fileObj.Close()
	//创建一个都对象
	reader := bufio.NewReader(fileObj)
	i :=0
	var chessMap2 [11][11]int
	for  {

		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n",line)
		line2 := strings.Trim(line, "\n")
		fmt.Printf("%#v\n",line2)
		fmt.Printf("%#v\n",len(line2))
		if err == io.EOF{
			fmt.Println("存盘数据读取完毕，原始数组已恢复...")
			break
		}
		if err != nil {
			fmt.Println("存盘数据读取失败...err",err)
			return
		}
		fmt.Println(line)
		getStr := strings.Split(line2,"\t")
		//fmt.Printf("%T",getStr[0])
		row,err := strconv.Atoi(getStr[0])
		fmt.Printf("%T %v\n",row,row)
		if  err != nil{
			fmt.Println("row Atoi() failed,err:",err)
			return
		}
		col,err := strconv.Atoi(getStr[1])
		fmt.Printf("%T %v\n",col,col)
		if  err != nil{
			fmt.Println("col Atoi() failed,err:",err)
			return
		}
		fmt.Println(len(getStr[1]))
		fmt.Println(len(getStr[2]))
		fmt.Printf("%#v\n",getStr[2][0])
		val,err := strconv.Atoi(getStr[2])
		fmt.Printf("%T %v\n",val,val)
		if  err != nil{
			fmt.Println("val Atoi() failed,err:",err)
			return
		}
		if i== 0{
			i++
			continue
		}
		chessMap2[row][col] = val
	}
	for _, v1 := range chessMap2 {
		for _, v2 := range v1 {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

}
