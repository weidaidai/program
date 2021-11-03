package main

import "fmt"

func main() {
	var op string
	fmt.Println("请输入你的操作：（y/n）")

	fmt.Scan(&op)
	fmt.Println("你的想法")
	fmt.Println("买一个西瓜")
	if op != "y" && op != "n" {
		goto END
	}
	fmt.Println("结束对话")
END:
}
