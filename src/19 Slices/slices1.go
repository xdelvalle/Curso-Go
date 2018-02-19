package main

import (
	"fmt"
)

func main() {
	// El Slice es como un array per dinamico

	// Declaracion
	var j []int
	fmt.Println("Slice j: ", j)

	x := []int{1, 2, 3}
	fmt.Println("Slice x: ", x)

	// Declarar usando 'make' indicando la longitud
	y := make([]int, 5)
	fmt.Println("Slice y: ", y)
	fmt.Println("Longitud y: ", len(y))
	fmt.Println("Capacidad y: ", cap(y))

	// Declarar usando 'make' indicando la longitud y capacidad
	k := make([]int, 5, 10)
	fmt.Println("Slice k: ", k)
	fmt.Println("Longitud k: ", len(k))
	fmt.Println("Capacidad k: ", cap(k))

	// Definimos un array con 10 elementos int
	var array = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array)

	var a []int
	var b []int

	fmt.Println("Slice a: ", a)
	fmt.Println("Slice b: ", b)

	// 'a' apunta el 3er y al 5to elemento del array 'array'
	a = array[2:5]
	// Ahora 'a' tiene los elementos array[2], array[3] y array[4]
	fmt.Println("Slice a: ", a)

	// 'b' es otro slice del array 'array'
	b = array[3:5]
	// Ahora 'b' tiene los elementos array[3] y array[4]
	fmt.Println("Slice b: ", b)

	// Los Slides apuntan al mismo array por lo que si se cambian en cualquier slide cambiará tb en el otro y tb en el array
	b[0] = 25
	fmt.Println("Asignamos b[0] = 25")
	fmt.Println("Slice b: ", b)
	fmt.Println("Slice a: ", a)
	fmt.Println("Array: ", array)
	fmt.Println("Cap(a) ", cap(a))
	fmt.Println("Cap(b) ", cap(b))
}
