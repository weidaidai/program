package main

import "fmt"

func test1(a int, b int, c string) (int, string, bool) {
	return a + b, c, true

}
func test2(a1, b2 int, c1 string) (ret int, str string, bl bool) {
	ret = a1 + b2
	str = c1
	bl = true
	return
	return a1 + b2, c1, true
}
func main() {
	v1, v2, _ := test1(10, 20, "hello")
	fmt.Println("v1:", v1, "v2", v2)
}
