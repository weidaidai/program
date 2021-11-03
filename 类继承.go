package main

import "fmt"

type mm struct {
	age   int
	name  string
	score float64
}

type ss struct {
	hh    mm
	shool string
}

func main() {
	s1 := ss{
		hh: mm{
			age:   48,
			score: 63,
		},
		shool: "人民大学",
	}
	fmt.Println(s1.hh.age)
	fmt.Println(s1.shool)
	type teacher struct {
		mm
		subject string
	}
	t1 :=teacher{}

}
