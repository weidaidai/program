package main

import (
	"fmt"
	"import/Sub"
)

//用struct来模拟类

//
type person struct {
	name  string
	age   int
	gen   string
	score float64
}

func (m *person) Eat() {
	fmt.Println("person is eating")
	fmt.Println(m.age - 78)
}
func main() {
	lili := person{
		name:  "韦可爱",
		age:   17,
		gen:   "men",
		score: 89.99,
	}
	fmt.Println(lili)
	lili.Eat()
	fmt.Println(lili)
	w := Sub.Sub(10,60)
	fmt.Println(w)
}
