package main

import "fmt"

func main() {

	//interface 可以用来处理多态还可以接受各种类型的数据
	//定义三个类型的接口
	var i, j, k interface{}
	name := []string{"lili", "keke"}
	i = name
	fmt.Println("i代表切片字符串", i)

	age := 20
	j = age
	fmt.Println("j代表int", j)
	gg := "kk"
	k = gg
	fmt.Println(k)
	//将interface当作一个函数的参数，，使用switch来判断用户输入的不同类型
	//创建有三个接口的切片
	array := make([]interface{}, 3)

	array[0] = "我是小可爱哦"
	array[1] = 123
	array[2] = true
	for _, value := range array {
		//v代表是array内的值，value.(type)为array的类型
		switch v := value.(type) {
		case int:
			fmt.Printf("当前数据为int %d\n", v)
		case string:
			fmt.Printf("当前数据为string %s\n", v)
		case bool:
			fmt.Printf("当前数据为bool %v\n", v)
		default:
			fmt.Printf("啥也不是")
		}
	}

}
