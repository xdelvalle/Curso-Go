package main

import (
	"fmt"
)

// Variadic functions. Funciones con el numero de parametros variables.
func main3() { // cambiar por main()
	fmt.Println(sumar(2, 3, 4, 5))
	printa("Hola", "Xavi", "que", "tal")
}

func sumar(numeros ...int) int {
	res := 0
	for _, numero := range numeros {
		res += numero
	}
	return res
}

func printa(cadena string, cadenas ...string) {
	for _, c := range cadenas {
		cadena += " " + c
	}

	fmt.Println(cadena)
}
