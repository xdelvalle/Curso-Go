package main

import (
	"fmt"
)

// Call stack, pila de llamadas. cola LIFO Last Input First Output
func main2() { // cambiar por main()
	fmt.Println("Entrando a main")
	f1()
	fmt.Println("Saliendo de main")
}

func f1() {
	fmt.Println("Entrando a f1")
	f2()
	fmt.Println("Saliendo de f1")
}

func f2() {
	fmt.Println("Entrando a f2")
	f3()
	fmt.Println("Saliendo de f2")
}

func f3() {
	fmt.Println("Entrando a f3")
	fmt.Println("Saliendo de f3")
}
