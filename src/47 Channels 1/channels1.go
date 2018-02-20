package main

import (
	"fmt"
	"time"
)

// Channels -> Comunicacion entre las gorutinas (concurrentes)

func main() {
	// Creamos el canal
	channel := make(chan string)

	// Llamamos a las funciones como gorutinas
	go enviarPing(channel)
	go imprimirPing(channel)

	// Escaneamos la entraa de datos para que no finalice main()
	var input string
	fmt.Scanln(&input)
	fmt.Println("Fin")
}

func enviarPing(c chan string) {
	for {
		// Enviamos valores al canal **Se queda parado hasta que el receptor reciba
		c <- "Ping"

		// No hace falta sleep porque el channel sincroniza el envio y la recepcion !!
	}
}

func imprimirPing(c chan string) {
	var contador int
	for {
		// Recibimos del canal
		contador++
		fmt.Println(<-c, " ", contador)
		time.Sleep(1 * time.Second)
	}
}
