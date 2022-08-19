package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{})  {
	v := reflect.TypeOf(x)
	fmt.Printf("type.name:%v  type.kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{})  {
	v := reflect.ValueOf(x)//返回值为x的实际值
	k := v.Kind()//获取v的种类
	fmt.Printf("the kind is %v\n", v)
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

//通过反射修改（设置）变量的值
func reflectSetValue1(x interface{})  {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)//修改的是副本， reflect包会引发panic
	}

}

//通过反射修改（设置）变量的值
func reflectSetValue2(x interface{})  {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)//修改的是副本， reflect包会引发panic
	}

}
func main() {
	var a float32 = 3.14
	reflectType(a)
	reflectValue(a)
	var b int64 = 100
	reflectType(b)
	reflectValue(b)
	//reflectSetValue1(b)修改系统报错
	reflectSetValue2(&b)
	fmt.Println(b)

	type cat struct {

	}
	var c = cat{}
	reflectType(c)


}
