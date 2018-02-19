package main

import (
	"fmt"
)

// Panic y recover. Panic se lanza cuando hay un error y para la ejecucion del programa, podemos lanzar nuestros propios panics

func main() {
	imprimir()
	fmt.Println("Hola Main")
}

func imprimir() {
	fmt.Println("Hola Xavi")

	// Para que no se pare el programa, hacemos el recover dentro de una funcion con defer
	defer func() {
		cadena := recover() // Recupera los datos y continua con la ejecución
		fmt.Println(cadena)
	}()

	panic("Error")
}
