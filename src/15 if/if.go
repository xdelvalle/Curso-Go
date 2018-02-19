package main

import (
	"fmt"
)

func main() {
	// Los if no llevan parentesis

	if 5 < 6 {
		fmt.Println("5 < 6")
	} else {
		fmt.Println("Pues no")
	}

	a := 6
	if a < 6 { // Permite ejecutar funciones antes de evaluar (dentro del if)
		fmt.Println("a < 6")
	} else if a > 6 {
		fmt.Println("a > 6")
	} else {
		fmt.Println("a = 6")
	}

	fmt.Println("Fin")
}
