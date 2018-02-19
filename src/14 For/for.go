package main

import (
	"fmt"
)

func main() {

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("*************************")

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
}
