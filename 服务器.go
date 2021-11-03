package main

import (
	"fmt"
	"net"
)

//创捷一个全局的client结构体

type client struct {
	name string
	addr string
	C    chan string
}

//创建一个全局变量在线列表

var onlinemap = make(map[string]client)

//创建全局的message

var massage = make(chan string)

func hsndleconnect(conn net.Conn) { //对接客户端数据通信

	defer conn.Close()
	//获取用户网络ip地址
	cltaddr := conn.RemoteAddr().String()
	//创建新链接的用户的结构体，默认用户是ip+端口
	clt := client{name: cltaddr, addr: cltaddr, C: make(chan string)}

	//添加到在线的用户map列表，key ：ip+端口 value：client
	onlinemap[cltaddr] = clt
	//发送用户上线给全局chan
	massage <- "{" + cltaddr + "}" + clt.name + "上线"
	func Massager()
	{
		//初始化在线列表
		onlinemap := make(map[string]client)
		//监听全局chan是否有数据,有数据放mag 无数据阻塞
		//循环从massage中读取
		for {
			mag := <-massage
			//循环发送消息给所有在线用户
			for _, clt := range onlinemap {
				clt.C
			}
		}
	}
	func main(){
		//启动服务器进行监听客户端的链接请求
		listener, err := net.Listen("tcp", "127.0.0.1:8080") //net包中有两个返回值是listerner监听器的作用,err
		if err != nil {
			fmt.Println("listen err:", err)
			return
		}

		defer listener.Close() //在程序结束后关闭
		//创建管理者go程
		go Massage()
		//循环监听客户链接请求
		for {
			conn, err := listener.Accept() //accept有两个返回值，当有用户链接上来的accept()会从阻塞状况解除返回一个conn用于通信接口
			if err != nil {
				fmt.Println("accept err:", err)

				continue
			}
			//启动新的go程和客户端进行通信
			go hsndleconnect(conn) //做参数传递到go程中
		}
	}

