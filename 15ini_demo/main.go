package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//ini配置文件解析器

//MySQLConfig MySQL配置结构体
type MysqlConfig struct {
	Address string `ini:"address"`
	Port int		`ini:"port"`
	Username string	`ini:"username"`
	Password string	`ini:"password"`
}
//RedisConfig redis配置结构体
type RedisConfig struct {
	Host string	`ini:"host"`
	Port int	`ini:"port"`
	Password string	`ini:"password"`
	Database int	`ini:"database"`
}
//Config 配置文件结构体
type Config struct {
	MysqlConfig	`ini:"mysql"`
	RedisConfig	`ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//0.参数检验
	t := reflect.TypeOf(data)
	//0.1 传递进来的data参数必须是指针类型
	fmt.Println(t,t.Kind())
	if t.Kind() != reflect.Ptr{
		err = fmt.Errorf("data(%v) should be a pointer", t.Kind())//格式化输出后返回一个err类型
		//err = errors.New("data should be a pointer")
		return
	}
	//0.2 传进来的data参数必须是结构体类型指针
	if t.Elem().Kind() != reflect.Struct{
		err = fmt.Errorf("data(%v) should be a struct", t.Elem().Kind())
		return
	}
	//1.读文件得到字节类型的数据
	//fileObj, err := os.OpenFile()
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b)//将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(b),"\r\n")
	fmt.Printf("%#v\n", lineSlice)
	//2.一行一行的读数据
	var structName string
	for index, line := range lineSlice {
		//去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		//跳过空行
		if len(line)==0{
			continue
		}
		//2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";")||strings.HasPrefix(line, "#"){
			continue
		}
		//2.2 如果是[]开头就表示是节(section)
		if strings.HasPrefix(line, "["){
			if line[0] != '[' || line[len(line)-1] !=']' {
				err = fmt.Errorf("line : %d syntax error", index+1)
				return
			}
			//把这一行首尾的[]去掉，取到中间的内容把首尾的空格去掉，拿到内容
			sectionName := strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName)==0 {
				err = fmt.Errorf("line:%d syntax errror", index+1)
				return
			}
			//根据sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini"){
					//说明找到了对应的嵌套结构体, 把字段名记录下来
					structName = field.Name//
					fmt.Printf("找到%s对应的嵌套结构体%s\n",sectionName,structName)
					break
				}
			}
		}else{
			//2.3 如果不是[]开头就是‘=’分割的键值对
			//2.3.1 分割key，value
			if strings.Index(line,"=") == -1 || strings.HasPrefix(line, "="){
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			idx := strings.Index(line, "=")
			key := strings.TrimSpace(line[:idx])
			value := strings.TrimSpace(line[idx+1:])
			//2.3.2 根据structName 去data里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)//拿到嵌套结构体的值信息
			sType := sValue.Type()//拿到嵌套结构体的类型信息
			//structObj := v.Elem().FieldByName(structName)
			if sType.Kind() != reflect.Struct{
				err = fmt.Errorf("data中的%s字段应该是一个结构体",structName)
				return
			}
			//2.3.3 遍历嵌套结构体的每一个字段，判断tag是否等于key
			var fieldName string
			var fieldType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fieldType = field
				if field.Tag.Get("ini") == key{
					//
					fieldName = field.Name
					fmt.Println(fieldName)
					break
				}
			}
			//2.3.4 如果 key=tag，给这个字段赋值
			//2.3.4.1 根据fieldName去取这个字段
			if len(fieldName)== 0{
				//在结构体中找不到对应的字段
				continue
			}
			fieldObj := sValue.FieldByName(fieldName)
			//2.3.4.2 进行赋值
			fmt.Println(fieldName,fieldType.Type.Kind())
			switch fieldType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value,10, 64)
				if err != nil {
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					return
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32,reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				fieldObj.SetFloat(valueFloat)
			}
 		}
	}
	return
}
func main() {
	var cfg Config
	err := loadIni("./conf.ini",&cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n",err)
		return
	}
	fmt.Println(cfg.MysqlConfig)
}
