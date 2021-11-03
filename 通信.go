package main

import (
	"fmt"
	"timllcc
)

var s = make(chan int)

func printer(stc string) {
	for _, v := range stc {
		fmt.Printf("%c", v) //打印字符
		time.Sleep(time.Millisecond * 300)
	}

}
func person1() {
	
	printer("hello")
	s <- 00
}
func person2() {
	<-s

	printer("world")

}
func main() {
	go person1()
	go person2()
	for {

	}
}
