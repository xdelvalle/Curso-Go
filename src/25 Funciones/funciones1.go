package main

import (
	"fmt"
)

func main1() { // cambiar por main()
	// Funciones

	imprimirNombre("Xavi")

	var res = suma(10, 5)
	fmt.Println("Resultado + : ", res)

	var res2 = resta(10, 5)
	fmt.Println("Resultado - : ", res2)

	// Firmas de las funciones
	fmt.Printf("Funcion suma: %T\n", suma)
	fmt.Printf("Funcion resta: %T\n", resta)
}

func imprimirNombre(nombre string) {
	fmt.Println("El nomnbre es: ", nombre)
}

func suma(x int, y int) int {
	return x + y
}

func resta(x, y int) (r int) {
	r = x - y
	return
}
