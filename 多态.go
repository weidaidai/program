package main

import "fmt"

//实现go多态，需要实现定义不同接口
//定义一个接口，类型为interface
type i interface {
	//接口函数可以有多个，但是这是能有函数原型，不可以实现
	attack()//定义一个函数接口

}
type humanlow struct {
	name string
	level int

}

func (a *humanlow)attack()  {//*humanlow指针值
	//让上面的attack 接收
	fmt.Println("我是",a.name,"等级",a.level)//a.level 为hunamlow中的
}

func main() {
		lowlevel:=humanlow{
			name:  "keke",
			level: 1,
		}
		lowlevel.attack()
}