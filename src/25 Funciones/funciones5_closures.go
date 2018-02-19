package main

import (
	"fmt"
)

func main5() { // cambiar por main()
	cadena := "Hola mundo"

	imprimir := print // print es una funcion
	imprimir(cadena)  // Ejecuta la funcion dentro de otra funcion

	// Otro modo de tener funciones dentro de funciones
	imprimir2 := func() {
		fmt.Println(cadena) // Tenemos acceso a las variables locales de main()
	}

	imprimir2()

	imprimir = print2
	imprimir("Hola Mundo 2")

	// imprimir = print3 // Da error porque no tienen la misma firma (los parametros son diferentes)
	// Estas son las firmas..
	fmt.Printf("Firma print: %T\n", print)
	fmt.Printf("Firma print2: %T\n", print2)
	fmt.Printf("Firma print3: %T\n", print3)

	// print es una funcion
	print4(print)

	// Las funciones son comparables con el valor nil
	var fb func()
	if fb == nil {
		fmt.Println("fb es nil")
	}

	inc := incrementar()
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())

	incremento()
	incremento()
	incremento()
	incremento()
}

func print(cadena string) {
	fmt.Println(cadena)
}

func print2(cadena string) {
	fmt.Println(cadena)
}

func print3(cadena1, cadena2 string) {
	fmt.Println(cadena1 + cadena2)
}

func print4(fprint func(string)) {
	fprint("Hola Mundo desde print4")
}

func incrementar() func() int {
	i := 0
	return func() (r int) {
		r = i
		i += 2
		return
	}
}

func incremento() {
	i := 0
	i++
	fmt.Println(i)
}
