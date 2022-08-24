package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request){
	//str, err := ioutil.ReadFile("./ex.html")
	str, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		fmt.Println("readFile failed, err:", err)
		w.Write([]byte("加载失败..."))
		return
	}
	w.Write(str)
}

func f2(w http.ResponseWriter, r *http.Request){
	//对于GET请求，参数都放在URL上（query param）,请求体中是没有数据的
	fmt.Println(r.URL)
	queryParam := r.URL.Query()//自动将URL中的参数识别整理成一个键值对形式的map并返回
	name := queryParam.Get("name")
	data := queryParam.Get("data")
	fmt.Println(name, data)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/xx/", f2)
	http.ListenAndServe("192.168.1.7:8000",nil)
}
