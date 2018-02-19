package main

import (
	"fmt"
)

func main() {
	// Bool 2 posibles valores, true o false

	var resultado bool

	resultado = (7 < 6)
	fmt.Println("5 < 6 = ", resultado)

	resultado = (5 < 6) && (3 > 1)
	fmt.Println("(5 < 6) && (3 > 1) = ", resultado)

	resultado = (5 < 1) || (3 > 6)
	fmt.Println("(5 < 6) || (3 > 6) = ", resultado)

	fmt.Println("!resultado = ", !resultado)
}
