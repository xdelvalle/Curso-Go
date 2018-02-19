package main

import (
	"fmt"
)

func main() {
	a := 25
	fmt.Println("Valor de a: ", a)
	fmt.Println("Direccion de a: ", &a)

	b := &a
	fmt.Println(b)  // Direccion
	fmt.Println(*b) // Valor

	// b es de tipo *int no int
	// b = 25 // ERROR

	*b = 32

	fmt.Println("Valor de a: ", a)

	a++

	fmt.Println("Valor que apunta b: ", *b)

	// El valor 0 de los punteros es nil
	if b != nil {
		fmt.Println("b != nil")
	}

	// Los punteros son comparables

	c := &a

	if b == c {
		fmt.Println("b y c son iguales")
	}

	// Obtener el puntero utilizando la funcion new()
	d := new(int) // *int
	fmt.Println("Direccion de d: ", d)
	fmt.Println("Valor de d:", *d)

	d = b
	fmt.Println("Direccion de d: ", d)
	fmt.Println("Valor de d:", *d)
	fmt.Println("Valor de d:", a)
	fmt.Println("Valor de d:", *b)
	fmt.Println("Valor de d:", *c)

	// Usar punteros en funciones

	// numero := 2
	// fmt.Println("Antes de incrementar: ", numero)
	// incrementar(numero)
	// fmt.Println("Despues de incrementar: ", numero)

	numero := 2
	fmt.Println("Antes de incrementar: ", numero)
	incrementar(&numero)
	fmt.Println("Despues de incrementar: ", numero)
}

func incrementar(numero *int) {
	*numero++
	fmt.Println("Numero en incrementar(): ", *numero)
}
