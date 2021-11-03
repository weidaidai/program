package main

import (
	"fmt"
	"time"
)

func main() {

	numchan := make(chan int)
	go func() {
		for i := 0; i < 50; i++ {
			data := <-numchan
			fmt.Println(data)
		}

	}()

	for i := 0; i < 50; i++ {
		numchan <- i
		fmt.Println("主go", i)

	}
	time.Sleep(5 * time.Second)

}
