package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiplicar(10))

	c1, c2, c3 := multiplicar(10)
	fmt.Println(c1, c2, c3)

	fmt.Println(multiplicar2(20))

	fmt.Println(retorno())

	fmt.Println(retorno2())
}

// Retorna las 3 variables n1, n2 y n3
func multiplicar(numero int) (n1, n2, n3 int) {
	n1 = numero * 10
	n2 = numero * 20
	n3 = numero * 30

	return
}

// Hay que indicar que retorna
func multiplicar2(numero int) (int, int, int) {
	n1 := numero * 10
	n2 := numero * 20
	n3 := numero * 30

	return n1, n2, n3
}

func retorno() (string, string) {
	return "Hola", "Mundo"
}

func retorno2() (string, int) {
	return "Hola", 5
}
