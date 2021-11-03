package main

import "fmt"

func test8() *string {
	city := "上海"
	ptr := &city

	return ptr

}
func main() {

	p1 := test8()
	fmt.Println(p1)

}
