package main

import (
	"fmt"
	"sync"
	"time"
)

var mutter sync.Mutex

func printer(stc string) {
	mutter.Lock()
	for _, v := range stc {
		fmt.Printf("%c", v) //打印字符
		time.Sleep(time.Millisecond * 300)
	}
mutter.Unlock()
}
func person1() {

	printer("hello")

}
func person2() {

	printer("world")

}
func main() {
	go person1()
	go person2()
	for {

	}
}
