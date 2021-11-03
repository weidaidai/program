package main

import (
	"fmt"
)

func main() {

	var i, j, k interface{}

	name := []string{"didi", "lili"}
	i = name
	fmt.Printf("%#v\n", i)
	age := 12
	j = age
	fmt.Println(j)
	str := "hello"
	k = str
	fmt.Println(k)
	vau, ok := k.(int)
	if !ok {
		fmt.Println("不是int")

	} else {
		fmt.Println("是int", vau)
	}
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "hello"
	array[2] = 3.45

	for _, vlue := range array {
		switch v := vlue.(type) {
		case int:
			fmt.Printf("当前为int：%#v\n", v)
		case string:
			fmt.Printf("当前为string：%#v\n", v)
		case bool:
			fmt.Printf("当前为bool：%#v\n", v)
		default:
			fmt.Println("都不是")
		}
	}

}
