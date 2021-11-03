package main

import "fmt"

func main() {
	var ss map[string]int
	fmt.Printf("%#v", ss)
	ss = map[string]int{"韦东启": 100}
	fmt.Println("韦东启")
	ss = make(map[string]int)
}
