package main

import (
	"fmt"
)

// Estructuras.

// Persona  Estructura Persona
type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	// Declaracion 1
	var p Persona
	p.Nombre = "Xavi"
	p.Edad = 38
	fmt.Println("Estructura p de tipo Persona: ", p)
	fmt.Println("Nombre p: ", p.Nombre)
	fmt.Println("Edad  p : ", p.Edad)

	// Declaracion 2
	p2 := Persona{Nombre: "Manolo", Edad: 30}
	fmt.Println("Nombre p2: ", p2.Nombre)
	fmt.Println("Edad p2  : ", p2.Edad)

	// Declaracion 3
	p3 := Persona{"Rita", 25}
	fmt.Println("Nombre p3: ", p3.Nombre)
	fmt.Println("Edad p3  : ", p3.Edad)

	// Inicializamos 3 varaiables de tipo Persona
	tom := Persona{"Tom", 60}
	bob := Persona{"Bob", 25}
	paul := Persona{"Paul", 43}

	tbOlder, tbDiff := Older(tom, bob)
	tpOlder, tpDiff := Older(tom, paul)
	bpOlder, bpDiff := Older(bob, paul)

	fmt.Printf("Entre %s y %s, %s es mayor por %d años\n", tom.Nombre, bob.Nombre, tbOlder.Nombre, tbDiff)
	fmt.Printf("Entre %s y %s, %s es mayor por %d años\n", tom.Nombre, paul.Nombre, tpOlder.Nombre, tpDiff)
	fmt.Printf("Entre %s y %s, %s es mayor por %d años\n", bob.Nombre, paul.Nombre, bpOlder.Nombre, bpDiff)
}

// Older Determina cual es mayor y la diferencia de edad
func Older(p1, p2 Persona) (Persona, int) {
	if p1.Edad > p2.Edad {
		return p1, p1.Edad - p2.Edad
	}

	return p2, p2.Edad - p1.Edad
}
