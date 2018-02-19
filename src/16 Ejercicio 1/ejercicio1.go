package main

import (
	"fmt"
)

func main() {
	// Contador de numeros impares

	encabezado := `
	***********************************************
	Contador de numeros impares
	***********************************************
	`

	fmt.Println(encabezado)

	fmt.Println("Insertar primer numero")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Println("Insertar segundo numero")
	var num2 int
	fmt.Scanln(&num2)

	var contador int

	for i := num1; i < num2; i++ {
		if i%2 != 0 {
			contador++
			fmt.Printf("El numero [%d] es impar\n", i)
		}
	}

	encabezado = `
	***********************************************
	Resultado
	***********************************************
	`
	fmt.Println(encabezado)
	fmt.Printf("Entre [%d] y [%d] hay [%d] numeros impares", num1, num2, contador)
}
