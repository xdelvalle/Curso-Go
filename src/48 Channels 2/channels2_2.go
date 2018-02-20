package main

import (
	"fmt"
	"time"
)

func main2() {
	numero := make(chan int)
	cuadrado := make(chan int)

	go func() {
		for x := 0; x < 5; x++ {
			numero <- x
		}
		close(numero)
	}()

	go func() {
		for {
			x, ok := <-numero // El segundo es un booleano que nos indica si el dato se ha recibido de un channel ABIERTO o CERRADO
			if !ok {
				break
			}
			cuadrado <- x * x
		}

		close(cuadrado)
	}()

	// Imprimimos en el main
	for {
		x, ok := <-cuadrado

		if !ok {
			break
		}

		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}

}
