package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
/*	resp, err := http.Get("http://192.168.1.7:8000/xx/?name=chenyi&data=2022-08-24")
	if err != nil {
		fmt.Println("get url failed, err:", err)
		return
	}
	defer resp.Body.Close()
	//从resp中把服务端返回的数据读出来
	//var data []byte
	//resp.Body.Read()
	//resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body failed, err:", err)
		return
	}
	fmt.Println(string(b))*/

	//自己构建一个http请求
	data := url.Values{}
	urlObj,_ := url.Parse("http://192.168.1.7:8000/xx/")
	data.Set("name","陈忆")
	data.Set("age", "25")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET",urlObj.String(),nil)
	if err != nil {
		fmt.Println("NewRequest err:", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("DefaultClient err:", err)
		return
	}
	fmt.Println(resp.Status)
}
