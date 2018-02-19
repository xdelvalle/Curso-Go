package main

import (
	"fmt"
)

func main() {
	var f32 float32     // 6 digitos decimales de precision
	var f64 float64     // 15 digitos decimales de precision
	var c64 complex64   // Numero complejo para float32
	var c128 complex128 // Numero complejo para float64

	fmt.Println("f32, f64, c64, c128 = ", f32, f64, c64, c128)

	f32 = .156
	f64 = .156

	fmt.Println("f32 * 58.565656 = ", f32*58.565656)
	fmt.Println("f64 * 58.565656 = ", f64*58.565656)
	fmt.Println("**************************************")

	c64 = complex(5, 6)
	fmt.Println("c64 = ", c64)
	fmt.Println("c64 * 85458.65 = ", c64*85458.65)

	c128 = complex(5, 6)
	fmt.Println("c128 = ", c128)
	fmt.Println("c128 * 85458.65 = ", c128*85458.65)

	fmt.Println("Real - c128 * 85458.65 = ", real(c128*85458.65))
	fmt.Println("Imaginaria - c128 * 85458.65 = ", imag(c128*85458.65))
}
