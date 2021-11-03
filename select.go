package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		for {
			fmt.Println("监听中")
			select {
			case dat1 := <-chan1:
				fmt.Println("从chan1读取数据", dat1)
			case dat2 := <-chan2:
				fmt.Println("从chan2读取数据", dat2)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			chan1 <- i
			time.Sleep(1 * time.Second / 2)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			chan2 <- i
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		fmt.Println("over")
		time.Sleep(5 * time.Second)
	}
}
