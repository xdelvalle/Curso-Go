package main

import (
	"fmt"
)

func main() {
	// Declaracion
	var x [3]int
	fmt.Println(x)

	var k [3][2]float64
	fmt.Println(k)

	// Asignar valores al array
	x[1] = 25
	fmt.Println(x)

	// Acceder a un indice del array
	fmt.Println(x[1])

	// Declarar e inicializar arrays
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	j := [...]int{1, 2, 3, 4}
	fmt.Println(j)

	f := [...]float64{1.2, 1.5, 8.5}
	fmt.Println(f)

	// Tamaño del array
	fmt.Println(len(f))

	// Comparar arrays
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 2, 4}
	//d := [4]int{1, 2, 3}

	fmt.Println(a == b)
	fmt.Println(a == c)
	//fmt.Println(a == d)
}
