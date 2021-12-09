package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getHandler(w http.ResponseWriter,r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	fmt.Println(r.Form)
	//遍历打印解析结果
	for key, value := range r.Form {
		fmt.Println(key, value)

	}
	w.Write([]byte("hello,world"))
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm)
	//遍历打印解析结果
	for key, value := range r.PostForm {
		fmt.Println(key, value)

	}
	// 2. 请求类型是application/json时从r.Body读取数据
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))

	w.Write([]byte("POST ok"))
}
func main(){
	http.HandleFunc("/post",postHandler)
	http.HandleFunc("/",getHandler)

   err:=http.ListenAndServe("127.0.0.1:8080",nil)
// 当handler为nil时，服务端调用http.DefaultServeMux进行处理
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
}
