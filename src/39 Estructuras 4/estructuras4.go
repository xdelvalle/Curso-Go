package main

import (
	"fmt"
)

// Punto struct
type Punto struct {
	x int
	y int
}

// Punto3D struct
type Punto3D struct {
	x int
	y int
	z int
	*Punto3D
}

// OpPunto tiene los metodos para realizar operaciones con puntos
type OpPunto struct {
}

func main() {
	// El valor 0 de una estructura es el valor 0 de cada uno de sus campos
	p := Punto{}
	fmt.Println("p: ", p)

	p2 := Punto3D{
		5,
		6,
		4,
		&Punto3D{
			6,
			4,
			6,
			nil,
		},
	}

	fmt.Println("p2: ", p2)

	// Comparando structs
	a := Punto{5, 6}
	b := Punto{7, 4}
	fmt.Println("a == b ?", a == b)

	c := Punto{7, 4}
	fmt.Println("b == c ?", b == c)

	// Utilizamos estructuras como indice dentro de un mapa
	figuras := make(map[Punto]string)
	figuras[a] = "Hola mundo"
	fmt.Println("figuras[a]: ", figuras[a])
}
