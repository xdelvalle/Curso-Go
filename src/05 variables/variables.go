package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Println(numero)

	numero = 40
	fmt.Println(numero)

	nombre := "Xavi"
	nombre = "Pepe"

	nombre, numero = "Xavi", 25
	fmt.Println(nombre, numero)

	nombre2 := "Ramon"
	fmt.Println(nombre, nombre2)

	nombre, nombre2 = nombre2, nombre
	fmt.Println(nombre, nombre2)

	nombre3, nombre := "Maria", "Juan"
	fmt.Println(nombre3, nombre)
}
