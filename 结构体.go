package main

import "fmt"

type teacher struct {
	name    string 
	age     int
	subject string
	gender  string
}

func main() {
	t1 := teacher{
		name:    "keke",
		age:     16,
		subject: "数学",

		gender: "男",
	}
	fmt.Println(t1)
}
