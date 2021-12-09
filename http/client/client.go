package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func get()  {
	//使用Get方法获取服务器响应包数据
	resp, err := http.Get("http://localhost:8080/?name=weidaidai&age=24&company=huaiguoshan")

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	// 获取服务器端读到的数据
	fmt.Println("Status = ", resp.Status)         // 状态
	fmt.Println("StatusCode = ", resp.StatusCode) // 状态码
	fmt.Println("Header = ", resp.Header)         // 响应头部
	fmt.Println("Body = ", resp.Body)             // 响应包体
	//读取body内的内容
	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}

//post方法

func post()  {
	url := "http://127.0.0.1:8080/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	// json
	contentType := "application/json"
	data := `{"name":"weidaidai","age":24}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}
func main() {
//get()
//post()
}

