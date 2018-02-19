package main

import (
	"fmt"
)

func main() {
	x := make([]byte, 4, 10)
	fmt.Println(x)

	x = []byte{'H', 'O', 'L', 'A'}
	fmt.Println(x)
	fmt.Printf("%q\n", x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%q\n", x[i])
	}

	// Asignamos espacio al slice
	//x[5] = ' ' // ERROR
	//fmt.Println(x)

	// Utilizando append
	x = append(x, ' ')
	fmt.Println(x)
	fmt.Println("Cap(x) ", cap(x)) // SI se pasa de capacidad, se hara un slice con el doble de capacidad (da 8)

	x = append(x, 'M', 'U', 'N', 'D', 'O')
	fmt.Println(x)
	fmt.Printf("%q\n", x)

	fmt.Println("Cap(x) ", cap(x)) // SI se pasa de capacidad, se hara un slice con el doble de capacidad (da 16)
	fmt.Println("Len(x) ", len(x))

	fmt.Println("************************************************************")

	var y []int

	for i := 1; i <= 15; i++ {
		y = append(y, i)

		fmt.Println("Slice y: ", y)
		fmt.Printf("Len: %d - Cap: %d - Elementos: %d\n", len(y), cap(y), i)
	}
}
