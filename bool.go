package main

import "fmt"

func main() {
	var zero bool

	boy := true

	gril := false

	fmt.Println(zero, boy, gril)
	fmt.Println(boy == zero)
	fmt.Println(boy != gril)
}
