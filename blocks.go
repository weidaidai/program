package main

import "fmt"

func main() {

	outer := 1
	{
		inner := 2
		fmt.Println(inner)
		fmt.Println(outer)
		{
			inner := 3
			fmt.Println(inner)
		}
	}

}
