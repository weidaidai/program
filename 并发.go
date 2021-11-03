package main

import (
	"fmt"
	"time"
)

// 这个用于子go

func display(num int) {
	count := 1
	// go程
	for {
		fmt.Println("这是子go程")
		count++

	}
}
func main() {

	for i := 0; i < 3; i++ {
		go display(i)
	}

	count := 1
	for {
		fmt.Println("这是主go程")
		count++
		time.Sleep(1 * time.Second)
	}
}
