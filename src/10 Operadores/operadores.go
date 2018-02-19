package main

import (
	"fmt"
)

func main() {

	// Operadores aritméticos

	// + Suma
	// - Resta
	// * Multiplicacion
	// / Division
	// % Modulo

	a := 5
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// Operaciones de comparacion
	// == Igual que
	// != Diferente que
	// > Mayor que
	// < Menor que
	// >= Mayor o igual que
	// <= Menor o igual que

	c := 3
	fmt.Println("************************************")

	fmt.Println("5 == 3 -> ", a == b)
	fmt.Println("3 == 3 -> ", b == c)
	fmt.Println("5 != 3 -> ", a != b)
	fmt.Println("3 != 3 -> ", c != b)
	fmt.Println("5 < 3 -> ", a < b)
	fmt.Println("3 <= 3 -> ", c <= b)
	fmt.Println("5 > 3 -> ", a > b)
	fmt.Println("3 >= 3 -> ", c >= b)
	fmt.Println("5 >= 3 -> ", a >= b)

	// Operadores de asignacion
	// += || a += b || a = a + b
	// -= || a -= b || a = a - b
	// *= || a *= b || a = a * b
	// /= || a /= b || a = a / b
	// %= || a %= b || a = a % b

	// Operadores logicos
	// && AND
	// || OR
	// ! NOT

	fmt.Println("************************************")
	fmt.Println("AND - &&")
	fmt.Println("true && true = ", true && true)
	fmt.Println("true && false = ", true && false)
	fmt.Println("false && false = ", false && false)
	// etc...

	// Operaciones de incremento/decremento
	// x++ -> x = x + 1
	// x-- -> x = x - 1
}
