package main

import (
	"errors"
	"fmt"
)

type Traveser func(ele interface{})

func Process(array interface{}, traveser Traveser) error {
	if array == nil{
		return errors.New("nil pointer")
	}
	var length int
	switch array.(type) {
	case []int:
		length = len(array.([]int))
	case []string:
		length = len(array.([]string))
	case []float32:
		length = len(array.([]float32))
	default:
		return errors.New("error type")
	}
	if length == 0{
		return errors.New("len is zero")
	}
	traveser(array)
	return nil
}

func SortByAscending(ele interface{})  {
	intSlice, ok := ele.([]int)
	if !ok {
		return
	}
	length := len(intSlice)

	for i := 0; i < length-1-i; i++ {
		isChange := false
		for j := 0; j < length-1-i; j++ {
			if intSlice[j] > intSlice[j+1]{
				isChange = true
				intSlice[j],intSlice[j+1] = intSlice[j+1], intSlice[j]
			}
		}
		if isChange == false{
			return
		}
	}
}
/*
具体操作：降序排序数组
 */

func SortByDescending(ele interface{})  {
	intSlice, ok := ele.([]int)
	if !ok {
		return
	}
	length := len(intSlice)
	for i := 0; i < length-1-i; i++ {
		isChange := false
		for j := 0; j < length-1-i; j++ {
			if intSlice[j] < intSlice[j+1] {
				isChange = true
				intSlice[j],intSlice[j+1] = intSlice[j+1], intSlice[j]
			}
		}
		if isChange == false{
			return
		}
	}
}
func main() {
	intSlice := make([]int, 0)
	intSlice = append(intSlice, 3, 1, 4, 2)

	Process(intSlice, SortByDescending)
	fmt.Println(intSlice)
	Process(intSlice, SortByAscending)
	fmt.Println(intSlice)

	stringSlice := make([]string, 0)
	stringSlice = append(stringSlice,"hello","world","china")

	Process(stringSlice, func(elem interface{}){
		if slice, ok := elem.([]string); ok{
			for index, value := range slice {
				fmt.Println("index:",index,"\tvalue:",value)
			}
		}
	})
	floatSlice := make([]float32,0)
	floatSlice = append(floatSlice,1.2,3.4,2.4)

	/*
	具体操作：使用匿名函数封装自定义操作
	 */
	Process(floatSlice, func(elem interface{}) {
		if slice, ok := elem.([]float32); ok{
			for index, value := range slice {
				slice[index] = value*2
			}
		}
	})
	fmt.Println(floatSlice)

}
