package main

import (
	"fmt"
	"time"
)

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	go generarNumeros(numero)

	go elevarAlCuadrado(numero, cuadrado)

	imprimir(cuadrado)
}

/*
CHANNELS DE ENTRADA O DE SALIDA
SALIDA: nombreFuncion(salida chan<- int)
ENTRADA: nombreFuncion(entrada <-chan int)
*/

func generarNumeros(out chan<- int) {
	for x := 0; x < 5; x++ {
		out <- x
	}
	close(out)
}

func elevarAlCuadrado(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x

		// Sin declarar si con channels de entrada o salida
		// in <- 25 // Peta en tiempo de EJECUCIÓN pq no se puede enviar datos en el channel de entrada. Para eso estan los undirectional channels

		// Declarando la direccionalidad de los channels
		// in <- 25 // Peta en tiempo de COMPILACIÓN
	}

	close(out)
}

func imprimir(in <-chan int) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
