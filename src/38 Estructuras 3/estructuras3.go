package main

import (
	"fmt"
	"math"
)

// Rectangulo struct
type Rectangulo struct {
	ancho float64
	alto  float64
}

// Circulo struct
type Circulo struct {
	radio float64
}

func main() {
	r1 := Rectangulo{12, 2}
	r2 := Rectangulo{9, 4}

	fmt.Println("Area R1: ", r1.area())
	fmt.Println("Area R2: ", r2.area())

	c1 := Circulo{8}
	c2 := Circulo{12}
	fmt.Println("Area C1: ", c1.area())
	fmt.Println("Area C2: ", c2.area())

	r1 = r1.inc(10)
	fmt.Println("R1: ", r1)

	r1.inc2(25) // No devolvemos nada, es un puntero y ya queda guardada r1
	fmt.Println("R1: ", r1)
}

func (r Rectangulo) area() float64 {
	return r.ancho * r.alto
}

func (c Circulo) area() float64 {
	return c.radio * c.radio * math.Pi
}

func (r Rectangulo) inc(i float64) Rectangulo {
	return Rectangulo{
		r.ancho * i,
		r.alto * i,
	}
}

func (r *Rectangulo) inc2(i float64) {
	r.ancho *= i
	r.alto *= i
}
